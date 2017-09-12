/*
 * Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
 *
 * This program and the accompanying materials are made available under
 * the terms of the under the Apache License, Version 2.0 (the "License”);
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package eureka

import (
	"fmt"

	"io"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/httpclient"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/serviceutil"
)

func Enable(cliConnection plugin.CliConnection, srInstanceName string, cfAppName string, authenticatedClient httpclient.AuthenticatedClient, instanceIndex *int, progressWriter io.Writer) (string, error) {
	return OperateOnApplication(cliConnection, srInstanceName, cfAppName, authenticatedClient, instanceIndex, progressWriter, serviceutil.ServiceInstanceURL,
		func(authClient httpclient.AuthenticatedClient, accessToken string, eurekaUrl string, eurekaAppName string, instanceId string) error {
			// Delete any registration status override and assume the new status is UP until proven otherwise.
			_, err := authClient.DoAuthenticatedDelete(fmt.Sprintf("%seureka/apps/%s/%s/status?value=UP", eurekaUrl, eurekaAppName, instanceId), accessToken)
			return err
		})
}
