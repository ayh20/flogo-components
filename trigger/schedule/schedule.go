package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

// HandlerSettings - Settings for the Handler
type HandlerSettings struct {
	StartInterval  string `md:"startDelay"`     // The start delay (ex. 1m, 1h, etc.), immediate if not specified
	RepeatInterval string `md:"repeatInterval"` // The repeat interval (ex. 1m, 1h, etc.), doesn't repeat if not specified
	StartDay       string `md:"startDay"`       // The day the process should run (ex. Monday, Tuesday, etc.)
	StartTime      string `md:"startTime"`      // The Time the process should run (ex. 08:30, 12:00, 21:00)
}

var triggerMd = trigger.NewMetadata(&HandlerSettings{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

// Factory - base struct factory
type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	return &Trigger{}, nil
}

// Trigger - Type def
type Trigger struct {
	timers   []*Job
	handlers []trigger.Handler
	logger   log.Logger
}

// Initialize - Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()

	return nil
}

// Start implements ext.Trigger.Start
func (t *Trigger) Start() error {

	handlers := t.handlers

	for _, handler := range handlers {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		//  We have overloaded the input parms .... if we have a day specified then the task should first
		// run then.

		if s.StartDay == "" {
			if s.RepeatInterval == "" {
				err = t.scheduleOnce(handler, s)
				if err != nil {
					return err
				}
			} else {
				err = t.scheduleRepeating(handler, s)
				if err != nil {
					return err
				}
			}
		} else {
			if s.RepeatInterval == "" {
				err = t.scheduleOnceOnDay(handler, s)
				if err != nil {
					return err
				}
			} else {
				err = t.scheduleRepeatingonDay(handler, s)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *Trigger) Stop() error {

	for _, timer := range t.timers {

		if timer.IsRunning() {
			timer.Quit <- true
		}
	}

	t.timers = nil

	return nil
}

func (t *Trigger) scheduleOnce(handler trigger.Handler, settings *HandlerSettings) error {
	t.logger.Info("Scheduling a Once timer")
	seconds := 0

	if settings.StartInterval != "" {
		d, err := time.ParseDuration(settings.StartInterval)
		if err != nil {
			return fmt.Errorf("unable to parse start delay: %s", err.Error())
		}

		seconds = int(d.Seconds())
		t.logger.Debugf("Scheduling action to run once in %d seconds", seconds)
	}

	var timerJob *Job

	fn := func() {
		t.logger.Info("Executing \"Once\" timer trigger")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Error("Error running handler: ", err.Error())
		}

		if timerJob != nil {
			timerJob.Quit <- true
		}
	}

	if seconds == 0 {
		t.logger.Info("Start delay not specified, executing action immediately")
		fn()
	} else {
		timerJob := Every(seconds).Seconds()
		timerJob, err := timerJob.NotImmediately().Run(fn)
		if err != nil {
			t.logger.Error("Error scheduling execute \"once\" timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	}

	return nil
}

func (t *Trigger) scheduleRepeating(handler trigger.Handler, settings *HandlerSettings) error {
	t.logger.Info("Scheduling a repeating timer")

	startSeconds := 0

	if settings.StartInterval != "" {
		d, err := time.ParseDuration(settings.StartInterval)
		if err != nil {
			return fmt.Errorf("unable to parse start delay: %s", err.Error())
		}

		startSeconds = int(d.Seconds())
		t.logger.Infof("Scheduling action to start in %d seconds", startSeconds)
	}

	d, err := time.ParseDuration(settings.RepeatInterval)
	if err != nil {
		return fmt.Errorf("unable to parse repeat interval: %s", err.Error())
	}

	repeatInterval := int(d.Seconds())
	t.logger.Infof("Scheduling action to repeat every %d seconds", repeatInterval)

	fn := func() {
		t.logger.Info("Executing \"Repeating\" timer")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Error("Error running handler: ", err.Error())
		}
	}

	if startSeconds == 0 {
		timerJob, err := Every(repeatInterval).Seconds().Run(fn)
		if err != nil {
			t.logger.Error("Error scheduling repeating timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	} else {

		timerJob := Every(startSeconds).Seconds()

		fn2 := func() {
			t.logger.Info("Executing first run of repeating timer")

			_, err := handler.Handle(context.Background(), nil)
			if err != nil {
				t.logger.Error("Error running handler: ", err.Error())
			}

			if timerJob != nil {
				timerJob.Quit <- true
			}

			timerJob, err := Every(repeatInterval).Seconds().NotImmediately().Run(fn)
			if err != nil {
				t.logger.Error("Error scheduling repeating timer: ", err.Error())
			}

			t.timers = append(t.timers, timerJob)
		}

		timerJob, err := timerJob.NotImmediately().Run(fn2)
		if err != nil {
			t.logger.Error("Error scheduling delayed start repeating timer: ", err.Error())
		}

		t.timers = append(t.timers, timerJob)
	}

	return nil
}
func (t *Trigger) scheduleOnceOnDay(handler trigger.Handler, settings *HandlerSettings) error {
	t.logger.Info("Scheduling a Once on a named Day/time timer")
	//ti := time.Now()
	//t.logger.Infof("Will be run @ %s %s now %s", settings.StartTime, settings.StartDay, ti.UTC())

	// StartInterval is ignored in this branch
	if settings.StartInterval != "" {
		return fmt.Errorf("StartDay and StartInterval not compatible")
	}

	var timerJob *Job
	var err error

	fn := func() {

		t.logger.Debug("Executing \"Once on Given day and time\" timer trigger")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Error("Error running handler: ", err.Error())
		}

		// remove job now that it's run once
		if timerJob != nil {
			timerJob.Quit <- true
		}
	}

	switch settings.StartDay {
	case "Sunday":
		timerJob, err = Every().Sunday().At(settings.StartTime).Run(fn)
	case "Monday":
		timerJob, err = Every().Monday().At(settings.StartTime).Run(fn)
	case "Tuesday":
		timerJob, err = Every().Tuesday().At(settings.StartTime).Run(fn)
	case "Wednesday":
		timerJob, err = Every().Wednesday().At(settings.StartTime).Run(fn)
	case "Thursday":
		timerJob, err = Every().Thursday().At(settings.StartTime).Run(fn)
	case "Friday":
		timerJob, err = Every().Friday().At(settings.StartTime).Run(fn)
	case "Saturday":
		timerJob, err = Every().Saturday().At(settings.StartTime).Run(fn)
	case "Everyday":
		timerJob, err = Every().Day().At(settings.StartTime).Run(fn)
	}

	if err != nil {
		t.logger.Error("Error scheduling execute \"Once on Given day and time\" timer: ", err.Error())
	}

	t.NextRun(timerJob)

	t.timers = append(t.timers, timerJob)

	return nil
}
func (t *Trigger) scheduleRepeatingonDay(handler trigger.Handler, settings *HandlerSettings) error {
	t.logger.Info("Scheduling a repeating on a named Day/time timer")

	// StartInterval is ignored in this branch
	if settings.StartInterval != "" {
		return fmt.Errorf("StartDay and StartInterval not compatible")
	}

	var timerJob *Job
	var err error

	fn := func() {
		t.logger.Debug("Executing \"Repeating on Given day and time\" schedule trigger")

		_, err := handler.Handle(context.Background(), nil)
		if err != nil {
			t.logger.Error("Error running handler: ", err.Error())
		}
		t.NextRun(timerJob)
	}

	switch settings.StartDay {
	case "Sunday":
		timerJob, err = Every().Sunday().At(settings.StartTime).Run(fn)
	case "Monday":
		timerJob, err = Every().Monday().At(settings.StartTime).Run(fn)
	case "Tuesday":
		timerJob, err = Every().Tuesday().At(settings.StartTime).Run(fn)
	case "Wednesday":
		timerJob, err = Every().Wednesday().At(settings.StartTime).Run(fn)
	case "Thursday":
		timerJob, err = Every().Thursday().At(settings.StartTime).Run(fn)
	case "Friday":
		timerJob, err = Every().Friday().At(settings.StartTime).Run(fn)
	case "Saturday":
		timerJob, err = Every().Saturday().At(settings.StartTime).Run(fn)
	case "Everyday":
		timerJob, err = Every().Day().At(settings.StartTime).Run(fn)
	}

	t.NextRun(timerJob)

	if err != nil {
		t.logger.Error("Error scheduling execute \"once\" timer: ", err.Error())
	}

	t.timers = append(t.timers, timerJob)

	return nil
}

// NextRun - When will task actually run !
func (t *Trigger) NextRun(job *Job) {

	actual, err := job.schedule.nextRun()
	if err != nil {
		t.logger.Error("Error determining nextRun: ", err.Error())
	}
	runTime := time.Now().Add(actual)
	//fmt.Printf("in xxx: Will run @ %v\r\n", runTime)
	t.logger.Infof("Task next scheduled for: %v", runTime)

}
