package grafana

import (
	"time"

	"github.com/aws/aws-sdk-go/service/managedgrafana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func waitWorkspaceCreated(conn *managedgrafana.ManagedGrafana, id string, timeout time.Duration) (*managedgrafana.WorkspaceDescription, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{managedgrafana.WorkspaceStatusCreating},
		Target:  []string{managedgrafana.WorkspaceStatusActive},
		Refresh: statusWorkspaceStatus(conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*managedgrafana.WorkspaceDescription); ok {
		return output, err
	}

	return nil, err
}

func waitWorkspaceUpdated(conn *managedgrafana.ManagedGrafana, id string, timeout time.Duration) (*managedgrafana.WorkspaceDescription, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{managedgrafana.WorkspaceStatusUpdating},
		Target:  []string{managedgrafana.WorkspaceStatusActive},
		Refresh: statusWorkspaceStatus(conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*managedgrafana.WorkspaceDescription); ok {
		return output, err
	}

	return nil, err
}

func waitWorkspaceDeleted(conn *managedgrafana.ManagedGrafana, id string, timeout time.Duration) (*managedgrafana.WorkspaceDescription, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{managedgrafana.WorkspaceStatusDeleting},
		Target:  []string{},
		Refresh: statusWorkspaceStatus(conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*managedgrafana.WorkspaceDescription); ok {
		return output, err
	}

	return nil, err
}
