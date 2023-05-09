package runners

// AutoResizeEnabledKey is the key of flag that enables pvc-autoresizer.
const AutoResizeEnabledKey = "resize.imaginekube.com/enabled"

// ResizeThresholdAnnotation is the key of resize threshold.
const ResizeThresholdAnnotation = "resize.imaginekube.com/threshold"

// ResizeInodesThresholdAnnotation is the key of resize threshold for inodes.
const ResizeInodesThresholdAnnotation = "resize.imaginekube.com/inodes-threshold"

// ResizeIncreaseAnnotation is the key of amount increased.
const ResizeIncreaseAnnotation = "resize.imaginekube.com/increase"

// StorageLimitAnnotation is the key of storage limit value
const StorageLimitAnnotation = "resize.imaginekube.com/storage-limit"

// PreviousCapacityBytesAnnotation is the key of previous volume capacity.
const PreviousCapacityBytesAnnotation = "resize.imaginekube.com/pre-capacity-bytes"

// AutoRestartEnabledKey  is the key of flag that enables pods-autoRestart.
const AutoRestartEnabledKey = "restart.imaginekube.com/enabled"

// SupportOnlineResize is the key of flag that the storage class support online expansion
const SupportOnlineResize = "restart.imaginekube.com/online-expansion-support"

// RestartSkip is the key of flag that the workload don't need autoRestart
const RestartSkip = "restart.imaginekube.com/skip"

// ResizingMaxTime is the key of flag that the maximum number of seconds that autoRestart can wait for pvc resize
const ResizingMaxTime = "restart.imaginekube.com/max-time"

// RestartStage is used to record whether autoRestart has finished shutting down the pod
const RestartStage = "restart.imaginekube.com/stage"

// RestartStopTime is used to record the time when the pod is closed
const RestartStopTime = "restart.imaginekube.com/stop-time"

// ExpectReplicaNums is used to record the value of replicas before restart
const ExpectReplicaNums = "restart.imaginekube.com/replica-nums"

// DefaultThreshold is the default value of ResizeThresholdAnnotation.
const DefaultThreshold = "10%"

// DefaultInodesThreshold is the default value of ResizeInodesThresholdAnnotation.
const DefaultInodesThreshold = "10%"

// DefaultIncrease is the default value of ResizeIncreaseAnnotation.
const DefaultIncrease = "10%"
