/*
Copyright 2019 Cloudera, Inc.  All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scheduler

import (
    "github.com/cloudera/yunikorn-core/pkg/cache"
    "github.com/cloudera/yunikorn-core/pkg/common/resources"
)

type SchedulingApplication struct {
    ApplicationInfo      *cache.ApplicationInfo
    Requests             *SchedulingRequests
    MayAllocatedResource *resources.Resource // Maybe allocated, set by scheduler

    // Private fields need protection
    queue *SchedulingQueue // queue the application is running in
}

func NewSchedulingApplication(appInfo *cache.ApplicationInfo) *SchedulingApplication {
    return &SchedulingApplication{
        ApplicationInfo: appInfo,
        Requests:        NewSchedulingRequests(),
    }
}
