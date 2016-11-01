package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/radanalyticsio/oshinko-rest/models"
)

// NewUpdateSingleClusterParams creates a new UpdateSingleClusterParams object
// with the default values initialized.
func NewUpdateSingleClusterParams() *UpdateSingleClusterParams {
	var ()
	return &UpdateSingleClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSingleClusterParamsWithTimeout creates a new UpdateSingleClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSingleClusterParamsWithTimeout(timeout time.Duration) *UpdateSingleClusterParams {
	var ()
	return &UpdateSingleClusterParams{

		timeout: timeout,
	}
}

/*UpdateSingleClusterParams contains all the parameters to send to the API endpoint
for the update single cluster operation typically these are written to a http.Request
*/
type UpdateSingleClusterParams struct {

	/*Cluster
	  Requested cluster update

	*/
	Cluster *models.NewCluster
	/*Name
	  Name of the cluster

	*/
	Name string

	timeout time.Duration
}

// WithCluster adds the cluster to the update single cluster params
func (o *UpdateSingleClusterParams) WithCluster(cluster *models.NewCluster) *UpdateSingleClusterParams {
	o.Cluster = cluster
	return o
}

// WithName adds the name to the update single cluster params
func (o *UpdateSingleClusterParams) WithName(name string) *UpdateSingleClusterParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSingleClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Cluster == nil {
		o.Cluster = new(models.NewCluster)
	}

	if err := r.SetBodyParam(o.Cluster); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}