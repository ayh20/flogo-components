package f1telemetry2021

// F1Header - Struct for the unpacking of the UDP Header
type F1Header struct {
	PacketFormat            uint16  `struc:"uint16,little"`  // 2019
	GameMajorVersion        uint8   `struc:"uint8,little"`   // Game major version - "X.00"
	GameMinorVersion        uint8   `struc:"uint8,little"`   // Game minor version - "1.XX"
	PacketVersion           uint8   `struc:"uint8,little"`   // Version of this packet type, all start from 1
	PacketID                uint8   `struc:"uint8,little"`   // Identifier for the packet type, see below
	SessionUID              uint64  `struc:"uint64,little"`  // Unique identifier for the session
	SessionTime             float32 `struc:"float32,little"` // Session timestamp
	FrameIdentifier         uint    `struc:"uint,little"`    // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   `struc:"uint8,little"`   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   `struc:"uint8,little"`   // Index of secondary player's car in the array (splitscreen) 255 if no second player

}

// F1CarMotion (Type 0 x22) - Struct for the unpacking of the UDP Motion format
// Frequency: Rate as specified in menus
type F1CarMotion struct {
	X          float32 `struc:"float32,little"` // World space position X
	Y          float32 `struc:"float32,little"` // World space position Y
	Z          float32 `struc:"float32,little"` // World space position Z
	Xv         float32 `struc:"float32,little"` // Velocity in world space X
	Yv         float32 `struc:"float32,little"` // Velocity in world space Y
	Zv         float32 `struc:"float32,little"` // Velocity in world space Z
	Xd         int16   `struc:"int16,little"`   // World space forward direction X
	Yd         int16   `struc:"int16,little"`   // World space forward direction Y
	Zd         int16   `struc:"int16,little"`   // World space forward direction Z
	Xr         int16   `struc:"int16,little"`   // World space right direction X
	Yr         int16   `struc:"int16,little"`   // World space right direction Y
	Zr         int16   `struc:"int16,little"`   // World space right direction Z
	Gforcelat  float32 `struc:"float32,little"` // Lateral G-Force component
	Gforcelon  float32 `struc:"float32,little"` // Longitudinal G-Force component
	Gforcevert float32 `struc:"float32,little"` // Vertical G-Force component
	Yaw        float32 `struc:"float32,little"` // Yaw angle in radians
	Pitch      float32 `struc:"float32,little"` // Pitch angle in radians
	Roll       float32 `struc:"float32,little"` // Roll angle in radians
}

// F1CarMotionExtra (Type 0 x1) - Struct for the unpacking of the UDP data format
// Frequency: Rate as specified in menus
type F1CarMotionExtra struct {
	SuspPosRL          float32 `struc:"float32,little"` // Suspension position RL, RR, FL, FR  F array
	SuspPosRR          float32 `struc:"float32,little"`
	SuspPosFL          float32 `struc:"float32,little"`
	SuspPosFR          float32 `struc:"float32,little"`
	SuspVelRL          float32 `struc:"float32,little"` // Suspension velocity RL, RR, FL, FR  F array
	SuspVelRR          float32 `struc:"float32,little"`
	SuspVelFL          float32 `struc:"float32,little"`
	SuspVelFR          float32 `struc:"float32,little"`
	SuspAccelerationRL float32 `struc:"float32,little"` // Suspension acceleration RL, RR, FL, FR
	SuspAccelerationRR float32 `struc:"float32,little"`
	SuspAccelerationFL float32 `struc:"float32,little"`
	SuspAccelerationFR float32 `struc:"float32,little"`
	WheelspeedRL       float32 `struc:"float32,little"` // Wheel Speed RL, RR, FL, FR  F array
	WheelspeedRR       float32 `struc:"float32,little"`
	WheelspeedFL       float32 `struc:"float32,little"`
	WheelspeedFR       float32 `struc:"float32,little"`
	WheelslipRL        float32 `struc:"float32,little"` // Wheel Slip RL, RR, FL, FR  F array
	WheelslipRR        float32 `struc:"float32,little"`
	WheelslipFL        float32 `struc:"float32,little"`
	WheelslipFR        float32 `struc:"float32,little"`
	XLocalVelocity     float32 `struc:"float32,little"` // Velocity in local space
	YLocalVelocity     float32 `struc:"float32,little"` // Velocity in local space
	ZLocalVelocity     float32 `struc:"float32,little"` // Velocity in local space
	Angvelx            float32 `struc:"float32,little"` // angular velocity x-component
	Angvely            float32 `struc:"float32,little"` // angular velocity y-component
	Angvelz            float32 `struc:"float32,little"` // angular velocity z-component
	AngAccX            float32 `struc:"float32,little"` // Angular acceleration x-component
	AngAccY            float32 `struc:"float32,little"` // Angular acceleration y-component
	AngAccZ            float32 `struc:"float32,little"` // Angular acceleration z-component
	FrontWheelsAngle   float32 `struc:"float32,little"` // Current front wheels angle in radians
}

type F1SessionZone struct {
	ZoneStart1 float32 `struc:"float32,little"` // Fraction (0..1) of way through the lap the marshal Zone starts
	ZoneFlag1  int8    `struc:"int8,little"`    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}
type F1SessionWeatherForecast struct {
	SessionType            uint8 `struc:"uint8,little"` // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2,  12 = Time Trial
	TimeOffset             uint8 `struc:"uint8,little"` // Time in minutes the forecast is for
	Weather                uint8 `struc:"uint8,little"` // Weather - 0 = clear, 1 = light cloud, 2 = overcast 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature       int8  `struc:"int8,little"`  // Track temp. in degrees Celsius
	TrackTemperatureChange int8  `struc:"int8,little"`  // Track temp. change – 0 = up, 1 = down, 2 = no change
	AirTemperature         int8  `struc:"int8,little"`  // Air temp. in degrees celsius
	AirTemperatureChange   int8  `struc:"int8,little"`  // Air temp. change – 0 = up, 1 = down, 2 = no change
	RainPercentage         uint8 `struc:"uint8,little"` // Rain percentage (0-100)

}

// F1Session (Type 1 x1) - Struct for the unpacking of the UDP data format
// Frequency: 2 per second
type F1Session struct {
	Weather             uint8  `struc:"uint8,little"`  // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature    int8   `struc:"int8,little"`   // Track temp. in degrees celsius
	AirTemperature      int8   `struc:"int8,little"`   // Air temp. in degrees celsius
	TotalLaps           uint8  `struc:"uint8,little"`  // Total number of laps in this race
	TrackLength         uint16 `struc:"uint16,little"` // Track length in metres
	SessionType         uint8  `struc:"uint8,little"`  // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ,  10 = R, 11 = R2, 12 = Time Trial
	TrackID             int8   `struc:"int8,little"`   // -1 for unknown, 0-21 for tracks, see appendix
	Formula             uint8  `struc:"uint8,little"`  // Formular, 0 = F1 modern, 1 = F1 classic, 2 - F2, 3 = F1 Generic
	SessionTimeLeft     uint16 `struc:"uint16,little"` // Time left in session in seconds
	SessionDuration     uint16 `struc:"uint16,little"` // Session duration in seconds
	PitSpeedLimit       uint8  `struc:"uint8,little"`  // Pit speed limit in kilometres per hour
	GamePaused          uint8  `struc:"uint8,little"`  // Whether the game is paused
	IsSpectating        uint8  `struc:"uint8,little"`  // Whether the player is spectating
	SpectatorCarIndex   uint8  `struc:"uint8,little"`  // Index of the car being spectated
	SliProNativeSupport uint8  `struc:"uint8,little"`  // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones     uint8  `struc:"uint8,little"`  // Number of marshal zones to follow
	// Zones Array
	MarshallZone    [21]F1SessionZone `struc:"[21]F1SessionZone"` // ,sizefrom=NumMarshalZones session zone
	SafetyCarStatus uint8             `struc:"uint8,little"`      // 0 = no safety car, 1 = full safety car 2 = virtual safety car
	NetworkGame     uint8             `struc:"uint8,little"`      // 0 = offline, 1 = online
	// Forecast array
	NumWeatherForecastSamples uint8                        `struc:"uint8,little"`                 // Number of weather samples to follow
	WeatherForecastSample     [56]F1SessionWeatherForecast `struc:"[56]F1SessionWeatherForecast"` // Array of weather forecast samples
	ForecastAccuracy          uint8                        `struc:"uint8,little"`                 // 0 = Perfect, 1 = Approximate
	//
	AIDifficulty           uint8  `struc:"uint8,little"`  // AI Difficulty rating – 0-110
	SeasonLinkIdentifier   uint32 `struc:"uint32,little"` // Identifier for season - persists across saves
	WeekendLinkIdentifier  uint32 `struc:"uint32,little"` // Identifier for weekend - persists across saves
	SessionLinkIdentifier  uint32 `struc:"uint32,little"` // Identifier for session - persists across saves
	PitStopWindowIdealLap  uint8  `struc:"uint8,little"`  // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap uint8  `struc:"uint8,little"`  // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition  uint8  `struc:"uint8,little"`  // Predicted position to rejoin at (player)
	SteeringAssist         uint8  `struc:"uint8,little"`  // 0 = off, 1 = on
	BrakingAssist          uint8  `struc:"uint8,little"`  // 0 = off, 1 = low, 2 = medium, 3 = high
	GearboxAssist          uint8  `struc:"uint8,little"`  // 1 = manual, 2 = manual & suggested gear, 3 = auto
	PitAssist              uint8  `struc:"uint8,little"`  // 0 = off, 1 = on
	PitReleaseAssist       uint8  `struc:"uint8,little"`  // 0 = off, 1 = on
	ERSAssist              uint8  `struc:"uint8,little"`  // 0 = off, 1 = on
	DRSAssist              uint8  `struc:"uint8,little"`  // 0 = off, 1 = on
	DynamicRacingLine      uint8  `struc:"uint8,little"`  // 0 = off, 1 = corners only, 2 = full
	DynamicRacingLineType  uint8  `struc:"uint8,little"`  // 0 = 2D, 1 = 3D

}

// F1LapData (Type 2 x22) - Struct for the unpacking of the UDP data format
// Frequency: Rate as specified in menus
type F1LapData struct {
	LastLapTime    uint32 `struc:"uint32,little"` // Last lap time in seconds
	CurrentLapTime uint32 `struc:"uint32,little"` // Current time around the lap in seconds
	//	BestLapTime       float32 `struc:"float32,little"` // Best lap time of the session in seconds
	Sector1Time                 uint16  `struc:"uint16,little"`  // Sector 1 time in seconds
	Sector2Time                 uint16  `struc:"uint16,little"`  // Sector 2 time in seconds
	LapDistance                 float32 `struc:"float32,little"` // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance               float32 `struc:"float32,little"` // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta              float32 `struc:"float32,little"` // Delta in seconds for safety car
	CarPosition                 uint8   `struc:"uint8,little"`   // Car race position
	CurrentLapNum               uint8   `struc:"uint8,little"`   // Current lap number
	PitStatus                   uint8   `struc:"uint8,little"`   // 0 = none, 1 = pitting, 2 = in pit area
	Sector                      uint8   `struc:"uint8,little"`   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid           uint8   `struc:"uint8,little"`   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                   uint8   `struc:"uint8,little"`   // Accumulated time penalties in seconds to be added
	Warnings                    uint8   `struc:"uint8,little"`   // Accumulated number of warnings issued
	NumUnservedDriveThroughPens uint8   `struc:"uint8,little"`   // Num drive through pens left to serve
	NumUnservedStopGoPens       uint8   `struc:"uint8,little"`   // Num stop go pens left to serve
	GridPosition                uint8   `struc:"uint8,little"`   // Grid position the vehicle started the race in
	DriverStatus                uint8   `struc:"uint8,little"`   // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus                uint8   `struc:"uint8,little"`   // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = disqualified, 5 = not classified, 6 = retired
	PitLaneTimerActive          uint8   `struc:"uint8,little"`   // Pit lane timing, 0 = inactive, 1 = active
	PitLaneTimeInLaneInMS       uint16  `struc:"uint16,little"`  // If active, the current time spent in the pit lane in ms
	PitStopTimerInMS            uint16  `struc:"uint16,little"`  // Time of the actual pit stop in ms
	PitStopShouldServePen       uint8   `struc:"uint8,little"`   // Whether the car should serve a penalty at this stop
}

// F1Event (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1Event struct {
	EventString string `struc:"[4]byte,little"` // Event string code
}

// F1EventDetailsFastestLap (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailsFastestLap struct {
	VehicleIndex uint8   `struc:"uint8,little"`   // Vehicle index of car achieving fastest lap
	LapTime      float32 `struc:"float32,little"` // Lap time is in seconds
}

// F1EventDetailsFastestLap (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailsPenalty struct {
	PenaltyType      uint8 `struc:"uint8,little"` // Penalty type – see Appendices
	InfringementType uint8 `struc:"uint8,little"` // Infringement type – see Appendices
	VehicleIdx       uint8 `struc:"uint8,little"` // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 `struc:"uint8,little"` // Vehicle index of the other car involved
	Time             uint8 `struc:"uint8,little"` // Time gained, or time spent doing action in seconds
	LapNum           uint8 `struc:"uint8,little"` // Lap the penalty occurred on
	PlacesGained     uint8 `struc:"uint8,little"` // Number of places gained by this
}

// F1EventDetailsFastestLap (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailsSpeedTrap struct {
	VehicleIndex            uint8   `struc:"uint8,little"`   // Vehicle index of the vehicle triggering speed trap
	Speed                   float32 `struc:"float32,little"` // Top speed achieved in kilometres per hour
	OverallFastestInSession uint8   `struc:"uint8,little"`   // Overall fastest speed in session = 1, otherwise 0
	DriverFastestInSession  uint8   `struc:"uint8,little"`   // Fastest speed for driver in session = 1, otherwise 0

}

// F1EventDetailStartLights (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailStartLights struct {
	NumLights uint8 `struc:"uint8,little"` // Vehicle index of car achieving fastest lap
}

// F1EventDetailsFlashback (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailsFlashback struct {
	FlashbackFrameIdentifier uint32  `struc:"uint32,little"`  // Frame identifier flashed back to
	FlashbackSessionTime     float32 `struc:"float32,little"` // Session time flashed back to
}

// F1EventDetailsButtons (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailsButtons struct {
	m_buttonStatus uint32 `struc:"uint32,little"` // Bit flags specifying which buttons are being pressed currently
}

// F1EventDetailsExtraIndex (Type 3 x1) - Struct for the unpacking of the UDP data format
// Frequency: When the event occurs
type F1EventDetailsExtraIndex struct {
	VehicleIndex uint8 `struc:"uint8,little"` // Vehicle index of related car
}

// F1Participant (Type 4 x1) - Struct for the unpacking of the UDP data format
// Frequency: Every 5 seconds
type F1Participant struct {
	NumActiveCars uint8 `struc:"uint8,little"` // Number of active cars in the data
}

// F1ParticipantData (Type 4 x22) - Struct for the unpacking of the UDP data format
// Frequency: Every 5 seconds
type F1ParticipantData struct {
	AiControlled  uint8  `struc:"uint8,little"`    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8  `struc:"uint8,little"`    // Driver id
	NetworkId     uint8  `struc:"uint8,little"`    // Network id – unique identifier for network players
	TeamID        uint8  `struc:"uint8,little"`    // Team id
	MyTeam        uint8  `struc:"uint8,little"`    // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8  `struc:"uint8,little"`    // Race number of the car
	Nationality   uint8  `struc:"uint8,little"`    // Nationality of the drive
	Name          string `struc:"[48]byte,little"` // Name of participant in UTF-8 format – null terminated.  Will be truncated with … (U+2026) if too long
	YourTelemetry uint8  `struc:"uint8,little"`    // The player's UDP setting, 0 = restricted, 1 = public
}

// F1SetupData (Type 5 x22)- Struct for the unpacking of the UDP data format
// Frequency: 2 per second
type F1SetupData struct {
	FrontWing              uint8   `struc:"uint8,little"`   // Front wing aero
	RearWing               uint8   `struc:"uint8,little"`   // Rear wing aero
	OnThrottle             uint8   `struc:"uint8,little"`   // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   `struc:"uint8,little"`   // Differential adjustment off throttle (percentage)
	FrontCamber            float32 `struc:"float32,little"` // Front camber angle (suspension geometry)
	RearCamber             float32 `struc:"float32,little"` // Rear camber angle (suspension geometry)
	FrontToe               float32 `struc:"float32,little"` // Front toe angle (suspension geometry)
	RearToe                float32 `struc:"float32,little"` // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   `struc:"uint8,little"`   // Front suspension
	RearSuspension         uint8   `struc:"uint8,little"`   // Rear suspension
	FrontAntiRollBar       uint8   `struc:"uint8,little"`   // Front anti-roll bar
	RearAntiRollBar        uint8   `struc:"uint8,little"`   // Front anti-roll bar
	FrontSuspensionHeight  uint8   `struc:"uint8,little"`   // Front ride height
	RearSuspensionHeight   uint8   `struc:"uint8,little"`   // Rear ride height
	BrakePressure          uint8   `struc:"uint8,little"`   // Brake pressure (percentage)
	BrakeBias              uint8   `struc:"uint8,little"`   // Brake bias (percentage)
	RearLeftTyrePressure   float32 `struc:"float32,little"` // Rear tyre pressure (PSI)
	RearRightTyrePressure  float32 `struc:"float32,little"` // Rear tyre pressure (PSI)
	FrontLeftTyrePressure  float32 `struc:"float32,little"` // Front tyre pressure (PSI)
	FrontRightTyrePressure float32 `struc:"float32,little"` // Front tyre pressure (PSI)
	Ballast                uint8   `struc:"uint8,little"`   // Ballast
	FuelLoad               float32 `struc:"float32,little"` // Fuel load
}

// F1CarTelemetryData (Type 6 x22) - Struct for the unpacking of the UDP data format
// Frequency: Rate as specified in menus
type F1CarTelemetryData struct {
	Speed                     uint16  `struc:"uint16,little"`  // Speed of car in kilometres per hour
	Throttle                  float32 `struc:"float32,little"` // Amount of throttle applied (0 to 100)
	Steer                     float32 `struc:"float32,little"` // Steering (-100 (full lock left) to 100 (full lock right))
	Brake                     float32 `struc:"float32,little"` // Amount of brake applied (0 to 100)
	Clutch                    uint8   `struc:"uint8,little"`   // Amount of clutch applied (0 to 100)
	Gear                      int8    `struc:"int8,little"`    // Gear selected (1-8, N=0, R=-1)
	EngineRPM                 uint16  `struc:"uint16,little"`  // Engine RPM
	Drs                       uint8   `struc:"uint8,little"`   // 0 = off, 1 = on
	RevLightsPercent          uint8   `struc:"uint8,little"`   // Rev lights indicator (percentage)
	RevLightsBitValue         uint16  `struc:"uint16,little"`  /// Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperatureRL       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureRR       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureFL       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureFR       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	TyresSurfaceTemperatureRL uint8   `struc:"uint8,little"`   // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureRR uint8   `struc:"uint8,little"`   // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureFL uint8   `struc:"uint8,little"`   // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureFR uint8   `struc:"uint8,little"`   // Tyres surface temperature (celsius)
	TyresInnerTemperatureRL   uint8   `struc:"uint8,little"`   // Tyres inner temperature (celsius)
	TyresInnerTemperatureRR   uint8   `struc:"uint8,little"`   // Tyres inner temperature (celsius)
	TyresInnerTemperatureFL   uint8   `struc:"uint8,little"`   // Tyres inner temperature (celsius)
	TyresInnerTemperatureFR   uint8   `struc:"uint8,little"`   // Tyres inner temperature (celsius)
	EngineTemperature         uint16  `struc:"uint16,little"`  // Engine temperature (celsius)
	TyresPressureRL           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	TyresPressureRR           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	TyresPressureFL           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	TyresPressureFR           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	SurfaceTypeRL             uint8   `struc:"uint8,little"`   // Driving surface
	SurfaceTypeRR             uint8   `struc:"uint8,little"`   // Driving surface
	SurfaceTypeFL             uint8   `struc:"uint8,little"`   // Driving surface
	SurfaceTypeFR             uint8   `struc:"uint8,little"`   // Driving surface
}

// F1CarTelemetryDataExtra - Struct for the unpacking of the UDP data format
// Frequency: Rate as specified in menus
type F1CarTelemetryDataExtra struct {
	MfdPanelIndex                uint8 `struc:"uint8,little"` // Index of MFD panel open - 255 = MFD closed, Single player, race – 0 = Car setup, 1 = Pits, 2 = Damage, 3 =  Engine, 4 = Temperatures (May vary depending on game mode)
	MfdPanelIndexSecondaryPlayer uint8 `struc:"uint8,little"` // See above
	SuggestedGear                int8  `struc:"int8,little"`  // Suggested gear for the player (1-8) 0 if no gear suggested

}

// F1CarStatus - (Type 7 x22) Struct for the unpacking of the UDP data format
// Frequency: Rate as specified in menus
type F1CarStatus struct {
	TractionControl         uint8   `struc:"uint8,little"`   // Traction control - 0 = off, 1 = medium, 2 = full
	AntiLockBrakes          uint8   `struc:"uint8,little"`   // 0 (off) - 1 (on)
	FuelMix                 uint8   `struc:"uint8,little"`   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8   `struc:"uint8,little"`   // Front brake bias (percentage)
	PitLimiterStatus        uint8   `struc:"uint8,little"`   // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32 `struc:"float32,little"` // Current fuel mass
	FuelCapacity            float32 `struc:"float32,little"` // Fuel capacity
	FuelRemainingLaps       float32 `struc:"float32,little"` // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                  uint16  `struc:"uint16,little"`  // Cars max RPM, point of rev limiter
	IdleRPM                 uint16  `struc:"uint16,little"`  // Cars idle RPM
	MaxGears                uint8   `struc:"uint8,little"`   // Maximum number of gears
	DrsAllowed              uint8   `struc:"uint8,little"`   // 0 = not allowed, 1 = allowed, -1 = unknown
	DrsActivationDistance   uint16  `struc:"uint16,little"`  // 0 = DRS not available, non-zero - DRS will be available in [X] metres
	ActualTyreCompound      uint8   `struc:"uint8,little"`   // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1, 7 = inter, 8 = wet, F1 Classic - 9 = dry, 10 = wet, F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard, 15 = wet
	VisualTyreCompound      uint8   `struc:"uint8,little"`   // F1 Visual - 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet - F1 Classic and F2 as above
	TyresAgeLaps            uint8   `struc:"uint8,little"`   // Age in laps of the current set of tyres
	VehicleFiaFlags         int8    `struc:"int8,little"`    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ErsStoreEnergy          float32 `struc:"float32,little"` // ERS energy store in Joules
	ErsDeployMode           uint8   `struc:"uint8,little"`   // ERS deployment mode, 0 = none, 1 = low, 2 = medium, 3 = high, 4 = overtake, 5 = hotlap
	ErsHarvestedThisLapMGUK float32 `struc:"float32,little"` // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32 `struc:"float32,little"` // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float32 `struc:"float32,little"` // ERS energy deployed this lap
	NetworkPaused           uint8   `struc:"uint8,little"`   // Whether the car is paused in a network game

}

// F1FinalClassificationData - (Type 8 x22) Struct for the unpacking of the UDP data format
// Frequency: Once at the end of a race
type F1FinalClassificationData struct {
	m_position     uint8 `struc:"uint8,little"` // Finishing position
	m_numLaps      uint8 `struc:"uint8,little"` // Number of laps completed
	m_gridPosition uint8 `struc:"uint8,little"` // Grid position of the car
	m_points       uint8 `struc:"uint8,little"` // Number of points scored
	m_numPitStops  uint8 `struc:"uint8,little"` // Number of pit stops made
	m_resultStatus uint8 `struc:"uint8,little"` // Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = didnotfinish, 5 = disqualified
	// 6 = not classified, 7 = retired
	m_bestLapTimeInMS  uint32   `struc:"uint32,little"`   // Best lap time of the session in milliseconds
	m_totalRaceTime    float64  `struc:"float64,little"`  // Total race time in seconds without penalties
	m_penaltiesTime    uint8    `struc:"uint8,little"`    // Total penalties accumulated in seconds
	m_numPenalties     uint8    `struc:"uint8,little"`    // Number of penalties applied to this driver
	m_numTyreStints    uint8    `struc:"uint8,little"`    // Number of tyres stints up to maximum
	m_tyreStintsActual [8]uint8 `struc:"[8]uint8,little"` // Actual tyres used by this driver
	m_tyreStintsVisual [8]uint8 `struc:"[8]uint8,little"` // Visual tyres used by this driver
}

// F1FinalClassificationPacket - (Type 8 x1) Struct for the unpacking of the UDP data format
// Frequency: Once at the end of a race
type F1FinalClassificationPacket struct {
	NumCars uint8 `struc:"uint8,little"` // Number of cars in the final classification
}

// F1LobbyInfoData - (Type 9 x22) Struct for the unpacking of the UDP data format
// Frequency: Two every second when in the lobby
// Lobby Info Packet - This packet details the players currently in a multiplayer lobby. It details each player’s selected car, any AI involved in the game and also the ready status of each of the participants.

type F1LobbyInfoData struct {
	AIControlled uint8  `struc:"uint8,little"`    // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId       uint8  `struc:"uint8,little"`    // Team id - see appendix (255 if no team currently selected)
	Nationality  uint8  `struc:"uint8,little"`    // Nationality of the driver
	Name         string `struc:"[48]byte,little"` // Name of participant in UTF-8 format – null terminated.  Will be truncated with ... (U+2026) if too long
	CarNumber    uint8  `struc:"uint8,little"`    // Car number of the player
	ReadyStatus  uint8  `struc:"uint8,little"`    // 0 = not ready, 1 = ready, 2 = spectating
}

// F1LobbyInfo - (Type 9 x1) Struct for the unpacking of the UDP data format
// Frequency: Two every second when in the lobby
type F1LobbyInfo struct {
	// Packet specific data
	m_numPlayers uint8 `struc:"uint8,little"` // Number of players in the lobby data
}

// F1CarDamageData - (Type 10 x22) Struct for the unpacking of the UDP data format
// Car Damage Packet - This packet details car damage parameters for all the cars in the race.
// Frequency: 2 per second
type F1CarDamageData struct {
	TyresWearRL          float32 `struc:"float32,little"` // Tyre wear percentage
	TyresWearRR          float32 `struc:"float32,little"` // Tyre wear percentage
	TyresWearFL          float32 `struc:"float32,little"` // Tyre wear percentage
	TyresWearFR          float32 `struc:"float32,little"` // Tyre wear percentage
	TyresDamageRL        uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageRR        uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageFL        uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageFR        uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	BrakesDamageRL       uint8   `struc:"uint8,little"`   // Brakes damage (percentage)
	BrakesDamageRR       uint8   `struc:"uint8,little"`   // Brakes damage (percentage)
	BrakesDamageFL       uint8   `struc:"uint8,little"`   // Brakes damage (percentage)
	BrakesDamageFR       uint8   `struc:"uint8,little"`   // Brakes damage (percentage)
	FrontLeftWingDamage  uint8   `struc:"uint8,little"`   // Front left wing damage (percentage)
	FrontRightWingDamage uint8   `struc:"uint8,little"`   // Front right wing damage (percentage)
	RearWingDamage       uint8   `struc:"uint8,little"`   // Rear wing damage (percentage)
	FloorDamage          uint8   `struc:"uint8,little"`   // Floor damage (percentage)
	DiffuserDamage       uint8   `struc:"uint8,little"`   // Diffuser damage (percentage)
	SidepodDamage        uint8   `struc:"uint8,little"`   // Sidepod damage (percentage)
	DrsFault             uint8   `struc:"uint8,little"`   // Indicator for DRS fault, 0 = OK, 1 = fault
	GearBoxDamage        uint8   `struc:"uint8,little"`   // Gear box damage (percentage)
	EngineDamage         uint8   `struc:"uint8,little"`   // Engine damage (percentage)
	EngineMGUHWear       uint8   `struc:"uint8,little"`   // Engine wear MGU-H (percentage)
	EngineESWear         uint8   `struc:"uint8,little"`   // Engine wear ES (percentage)
	EngineCEWear         uint8   `struc:"uint8,little"`   // Engine wear CE (percentage)
	EngineICEWear        uint8   `struc:"uint8,little"`   // Engine wear ICE (percentage)
	EngineMGUKWear       uint8   `struc:"uint8,little"`   // Engine wear MGU-K (percentage)
	EngineTCWear         uint8   `struc:"uint8,little"`   // Engine wear TC (percentage)
}

// F1SessionHistoryData - (Type 11 x1 x100 x8) Struct for the unpacking of the UDP data format
// Session History Packet
//
// This packet contains lap times and tyre usage for the session. This packet works slightly differently to other packets.
// To reduce CPU and bandwidth, each packet relates to a specific vehicle and is sent every 1/20 s,
// and the vehicle being sent is cycled through. Therefore in a 20 car race you should receive an update for each vehicle at least once per second.
// Note that at the end of the race, after the final classification packet has been sent,
// a final bulk update of all the session histories for the vehicles in that session will be sent.
// Frequency: 20 per second but cycling through cars
type F1SessionHistoryData struct {
	m_carIdx             uint8                     `struc:"uint8,little"`          // Index of the car this lap data relates to
	m_numLaps            uint8                     `struc:"uint8,little"`          // Num laps in the data (including current partial lap)
	m_numTyreStints      uint8                     `struc:"uint8,little"`          // Number of tyre stints in the data
	m_bestLapTimeLapNum  uint8                     `struc:"uint8,little"`          // Lap the best lap time was achieved on
	m_bestSector1LapNum  uint8                     `struc:"uint8,little"`          // Lap the best Sector 1 time was achieved on
	m_bestSector2LapNum  uint8                     `struc:"uint8,little"`          // Lap the best Sector 2 time was achieved on
	m_bestSector3LapNum  uint8                     `struc:"uint8,little"`          // Lap the best Sector 3 time was achieved on
	LapHistoryData       [100]F1LapHistoryData     `struc:"[100]F1LapHistoryData"` // 100 laps of data max
	TyreStintHistoryData [8]F1TyreStintHistoryData `struc:"[8]F1TyreStintHistoryData"`
}

type F1LapHistoryData struct {
	m_lapTimeInMS      uint32 `struc:"uint32,little"` // Lap time in milliseconds
	m_sector1TimeInMS  uint16 `struc:"uint16,little"` // Sector 1 time in milliseconds
	m_sector2TimeInMS  uint16 `struc:"uint16,little"` // Sector 2 time in milliseconds
	m_sector3TimeInMS  uint16 `struc:"uint16,little"` // Sector 3 time in milliseconds
	m_lapValidBitFlags uint8  `struc:"uint8,little"`  // 0x01 bit set-lap valid, 0x02 bit set-sector 1 valid, 0x04 bit set-sector 2 valid, 0x08 bit set-sector 3 valid
}

type F1TyreStintHistoryData struct {
	m_endLap             uint8 `struc:"uint8,little"` // Lap the tyre usage ends on (255 of current tyre)
	m_tyreActualCompound uint8 `struc:"uint8,little"` // Actual tyres used by this driver
	m_tyreVisualCompound uint8 `struc:"uint8,little"` // Visual tyres used by this driver
}
