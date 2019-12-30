/*
Copyright 2018 The Kubernetes Authors.

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

package main

import (
	"context"
	"flag"
	"github.com/rexray/gocsi"
	"k8s.io/klog"

	"ics-csi-driver/pkg/csi/provider"
	"ics-csi-driver/pkg/csi/service"
)

// main is ignored when this package is built as a go plug-in.
func main() {
	klog.InitFlags(nil)
	flag.Parse()
	gocsi.Run(
		context.Background(),
		service.Name,
		"A CSI plugin for Inspur ICSphere storage",
		usage,
		provider.New())
}

const usage = `    ICSPHERE_CSI_CONFIG
        Specifies the path to the csi-icsphere.conf file

        The default value is "/etc/cloud/csi-icsphere.conf"
`
