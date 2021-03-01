package f1telemetry2019proto

// F1Header - Struct for the unpacking of the UDP Header
type F1Header struct {
	PacketFormat     uint16  `struc:"uint16,little"`  // 2019
	GameMajorVersion uint8   `struc:"uint8,little"`   // Game major version - "X.00"
	GameMinorVersion uint8   `struc:"uint8,little"`   // Game minor version - "1.XX"
	PacketVersion    uint8   `struc:"uint8,little"`   // Version of this packet type, all start from 1
	PacketID         uint8   `struc:"uint8,little"`   // Identifier for the packet type, see below
	SessionUID       uint64  `struc:"uint64,little"`  // Unique identifier for the session
	SessionTime      float32 `struc:"float32,little"` // Session timestamp
	FrameIdentifier  uint    `struc:"uint,little"`    // Identifier for the frame the data was retrieved on
	PlayerCarIndex   uint8   `struc:"uint8,little"`   // Index of player's car in the array
}

// F1CarMotion (Type 0 x20) - Struct for the unpacking of the UDP Motion format
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

// F1Session (Type 1 x1) - Struct for the unpacking of the UDP data format
type F1Session struct {
	Weather             uint8   `struc:"uint8,little"`   // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature    int8    `struc:"int8,little"`    // Track temp. in degrees celsius
	AirTemperature      int8    `struc:"int8,little"`    // Air temp. in degrees celsius
	TotalLaps           uint8   `struc:"uint8,little"`   // Total number of laps in this race
	TrackLength         uint16  `struc:"uint16,little"`  // Track length in metres
	SessionType         uint8   `struc:"uint8,little"`   // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ,  10 = R, 11 = R2, 12 = Time Trial
	TrackID             int8    `struc:"int8,little"`    // -1 for unknown, 0-21 for tracks, see appendix
	Formula             uint8   `struc:"uint8,little"`   // Formular, 0 = F1 modern, 1 = F1 classic, 2 - F2, 3 = F1 Generic
	SessionTimeLeft     uint16  `struc:"uint16,little"`  // Time left in session in seconds
	SessionDuration     uint16  `struc:"uint16,little"`  // Session duration in seconds
	PitSpeedLimit       uint8   `struc:"uint8,little"`   // Pit speed limit in kilometres per hour
	GamePaused          uint8   `struc:"uint8,little"`   // Whether the game is paused
	IsSpectating        uint8   `struc:"uint8,little"`   // Whether the player is spectating
	SpectatorCarIndex   uint8   `struc:"uint8,little"`   // Index of the car being spectated
	SliProNativeSupport uint8   `struc:"uint8,little"`   // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones     uint8   `struc:"uint8,little"`   // Number of marshal zones to follow
	ZoneStart1          float32 `struc:"float32,little"` // Fraction (0..1) of way through the lap the marshal Zone starts
	ZoneFlag1           int8    `struc:"int8,little"`    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ZoneStart2          float32 `struc:"float32,little"` // Note defining as separate fields for the moment to make the code simpler
	ZoneFlag2           int8    `struc:"int8,little"`    // review soon....
	ZoneStart3          float32 `struc:"float32,little"`
	ZoneFlag3           int8    `struc:"int8,little"`
	ZoneStart4          float32 `struc:"float32,little"`
	ZoneFlag4           int8    `struc:"int8,little"`
	ZoneStart5          float32 `struc:"float32,little"`
	ZoneFlag5           int8    `struc:"int8,little"`
	ZoneStart6          float32 `struc:"float32,little"`
	ZoneFlag6           int8    `struc:"int8,little"`
	ZoneStart7          float32 `struc:"float32,little"`
	ZoneFlag7           int8    `struc:"int8,little"`
	ZoneStart8          float32 `struc:"float32,little"`
	ZoneFlag8           int8    `struc:"int8,little"`
	ZoneStart9          float32 `struc:"float32,little"`
	ZoneFlag9           int8    `struc:"int8,little"`
	ZoneStart10         float32 `struc:"float32,little"`
	ZoneFlag10          int8    `struc:"int8,little"`
	ZoneStart11         float32 `struc:"float32,little"`
	ZoneFlag11          int8    `struc:"int8,little"`
	ZoneStart12         float32 `struc:"float32,little"`
	ZoneFlag12          int8    `struc:"int8,little"`
	ZoneStart13         float32 `struc:"float32,little"`
	ZoneFlag13          int8    `struc:"int8,little"`
	ZoneStart14         float32 `struc:"float32,little"`
	ZoneFlag14          int8    `struc:"int8,little"`
	ZoneStart15         float32 `struc:"float32,little"`
	ZoneFlag15          int8    `struc:"int8,little"`
	ZoneStart16         float32 `struc:"float32,little"`
	ZoneFlag16          int8    `struc:"int8,little"`
	ZoneStart17         float32 `struc:"float32,little"`
	ZoneFlag17          int8    `struc:"int8,little"`
	ZoneStart18         float32 `struc:"float32,little"`
	ZoneFlag18          int8    `struc:"int8,little"`
	ZoneStart19         float32 `struc:"float32,little"`
	ZoneFlag19          int8    `struc:"int8,little"`
	ZoneStart20         float32 `struc:"float32,little"`
	ZoneFlag20          int8    `struc:"int8,little"`
	ZoneStart21         float32 `struc:"float32,little"`
	ZoneFlag21          int8    `struc:"int8,little"`
	SafetyCarStatus     uint8   `struc:"uint8,little"` // 0 = no safety car, 1 = full safety car 2 = virtual safety car
	NetworkGame         uint8   `struc:"uint8,little"` // 0 = offline, 1 = online
}

// F1LapData (Type 2 x20) - Struct for the unpacking of the UDP data format
type F1LapData struct {
	LastLapTime       float32 `struc:"float32,little"` // Last lap time in seconds
	CurrentLapTime    float32 `struc:"float32,little"` // Current time around the lap in seconds
	BestLapTime       float32 `struc:"float32,little"` // Best lap time of the session in seconds
	Sector1Time       float32 `struc:"float32,little"` // Sector 1 time in seconds
	Sector2Time       float32 `struc:"float32,little"` // Sector 2 time in seconds
	LapDistance       float32 `struc:"float32,little"` // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance     float32 `struc:"float32,little"` // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta    float32 `struc:"float32,little"` // Delta in seconds for safety car
	CarPosition       uint8   `struc:"uint8,little"`   // Car race position
	CurrentLapNum     uint8   `struc:"uint8,little"`   // Current lap number
	PitStatus         uint8   `struc:"uint8,little"`   // 0 = none, 1 = pitting, 2 = in pit area
	Sector            uint8   `struc:"uint8,little"`   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid uint8   `struc:"uint8,little"`   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties         uint8   `struc:"uint8,little"`   // Accumulated time penalties in seconds to be added
	GridPosition      uint8   `struc:"uint8,little"`   // Grid position the vehicle started the race in
	DriverStatus      uint8   `struc:"uint8,little"`   // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus      uint8   `struc:"uint8,little"`   // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = disqualified, 5 = not classified, 6 = retired
}

// F1Event (Type 3 x1) - Struct for the unpacking of the UDP data format
type F1Event struct {
	EventString string `struc:"[4]byte,little"` // Event string code
}

// F1EventDetailsFastestLap (Type 3 x1) - Struct for the unpacking of the UDP data format
type F1EventDetailsFastestLap struct {
	VehicleIndex uint8   `struc:"uint8,little"`   // Vehicle index of car achieving fastest lap
	LapTime      float32 `struc:"float32,little"` // Lap time is in seconds
}

// F1EventDetailsExtraIndex (Type 3 x1) - Struct for the unpacking of the UDP data format
type F1EventDetailsExtraIndex struct {
	VehicleIndex uint8 `struc:"uint8,little"` // Vehicle index of related car
}

// F1Participant (Type 4 x1) - Struct for the unpacking of the UDP data format
type F1Participant struct {
	NumActiveCars uint8 `struc:"uint8,little"` // Number of active cars in the data
}

// F1ParticipantData (Type 4 x20) - Struct for the unpacking of the UDP data format
type F1ParticipantData struct {
	AiControlled  uint8  `struc:"uint8,little"`    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8  `struc:"uint8,little"`    // Driver id
	TeamID        uint8  `struc:"uint8,little"`    // Team id
	RaceNumber    uint8  `struc:"uint8,little"`    // Race number of the car
	Nationality   uint8  `struc:"uint8,little"`    // Nationality of the drive
	Name          string `struc:"[48]byte,little"` // Name of participant in UTF-8 format – null terminated.  Will be truncated with … (U+2026) if too long
	YourTelemetry uint8  `struc:"uint8,little"`    // The player's UDP setting, 0 = restricted, 1 = public
}

// F1SetupData (Type 5 x20)- Struct for the unpacking of the UDP data format
type F1SetupData struct {
	FrontWing             uint8   `struc:"uint8,little"`   // Front wing aero
	RearWing              uint8   `struc:"uint8,little"`   // Rear wing aero
	OnThrottle            uint8   `struc:"uint8,little"`   // Differential adjustment on throttle (percentage)
	OffThrottle           uint8   `struc:"uint8,little"`   // Differential adjustment off throttle (percentage)
	FrontCamber           float32 `struc:"float32,little"` // Front camber angle (suspension geometry)
	RearCamber            float32 `struc:"float32,little"` // Rear camber angle (suspension geometry)
	FrontToe              float32 `struc:"float32,little"` // Front toe angle (suspension geometry)
	RearToe               float32 `struc:"float32,little"` // Rear toe angle (suspension geometry)
	FrontSuspension       uint8   `struc:"uint8,little"`   // Front suspension
	RearSuspension        uint8   `struc:"uint8,little"`   // Rear suspension
	FrontAntiRollBar      uint8   `struc:"uint8,little"`   // Front anti-roll bar
	RearAntiRollBar       uint8   `struc:"uint8,little"`   // Front anti-roll bar
	FrontSuspensionHeight uint8   `struc:"uint8,little"`   // Front ride height
	RearSuspensionHeight  uint8   `struc:"uint8,little"`   // Rear ride height
	BrakePressure         uint8   `struc:"uint8,little"`   // Brake pressure (percentage)
	BrakeBias             uint8   `struc:"uint8,little"`   // Brake bias (percentage)
	FrontTyrePressure     float32 `struc:"float32,little"` // Front tyre pressure (PSI)
	RearTyrePressure      float32 `struc:"float32,little"` // Rear tyre pressure (PSI)
	Ballast               uint8   `struc:"uint8,little"`   // Ballast
	FuelLoad              float32 `struc:"float32,little"` // Fuel load
}

// F1CarTelemetryData (Type 6 x20) - Struct for the unpacking of the UDP data format
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
	BrakesTemperatureRL       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureRR       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureFL       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureFR       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	TyresSurfaceTemperatureRL uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureRR uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureFL uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureFR uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresInnerTemperatureRL   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	TyresInnerTemperatureRR   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	TyresInnerTemperatureFL   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	TyresInnerTemperatureFR   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
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
type F1CarTelemetryDataExtra struct {
	ButtonStatus uint32 `struc:"uint32,little"` // Bit flags specifying which buttons are being pressed currently
}

// F1CarStatus - (Type 7 x20) Struct for the unpacking of the UDP data format
type F1CarStatus struct {
	TractionControl         uint8   `struc:"uint8,little"`   // 0 (off) - 2 (high)
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
	TyresWearRL             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyresWearRR             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyresWearFL             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyresWearFR             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	ActualTyreCompound      uint8   `struc:"uint8,little"`   // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1, 7 = inter, 8 = wet, F1 Classic - 9 = dry, 10 = wet, F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard, 15 = wet
	VisualTyreCompound      uint8   `struc:"uint8,little"`   // F1 Visual - 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet - Classic and F2 as above
	TyresDamageRL           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageRR           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageFL           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageFR           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	FrontLeftWingDamage     uint8   `struc:"uint8,little"`   // Front left wing damage (percentage)
	FrontRightWingDamage    uint8   `struc:"uint8,little"`   // Front right wing damage (percentage)
	RearWingDamage          uint8   `struc:"uint8,little"`   // Rear wing damage (percentage)
	EngineDamage            uint8   `struc:"uint8,little"`   // Engine damage (percentage)
	GearBoxDamage           uint8   `struc:"uint8,little"`   // Gear box damage (percentage)
	VehicleFiaFlags         uint8   `struc:"uint8,little"`   // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ErsStoreEnergy          float32 `struc:"float32,little"` // ERS energy store in Joules
	ErsDeployMode           uint8   `struc:"uint8,little"`   // ERS deployment mode, 0 = none, 1 = low, 2 = medium, 3 = high, 4 = overtake, 5 = hotlap
	ErsHarvestedThisLapMGUK float32 `struc:"float32,little"` // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32 `struc:"float32,little"` // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float32 `struc:"float32,little"` // ERS energy deployed this lap
}
