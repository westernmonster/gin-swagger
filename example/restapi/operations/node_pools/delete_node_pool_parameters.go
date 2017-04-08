package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// BindDeleteNodePool validates and binds request parameters to the gin
// context.
func BindDeleteNodePool(ctx *gin.Context) {
	params := &DeleteNodePoolParams{}
	err := params.bindRequest(ctx)
	if err != nil {
		errObj := err.(*errors.CompositeError)
		ctx.JSON(int(errObj.Code()), errObj)
		return
	}
	ctx.Set("params", params)
	ctx.Next()
}

// DeleteNodePoolParams contains all the bound params for the delete node pool operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteNodePool
type DeleteNodePoolParams struct {

	/*ID of the cluster.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-:]*[a-z0-9]$
	  In: path
	*/
	ClusterID string
	/*Name of the node pool.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-]*[a-z0-9]$
	  In: path
	*/
	NodePoolName string
}

// DeleteNodePoolParamsFromCtx gets the params struct from the gin context.
func DeleteNodePoolParamsFromCtx(ctx *gin.Context) *DeleteNodePoolParams {
	params, _ := ctx.Get("params")
	return params.(*DeleteNodePoolParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteNodePoolParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rClusterID := []string{ctx.Param("cluster_id")}
	if err := o.bindClusterID(rClusterID, true, formats); err != nil {
		res = append(res, err)
	}

	rNodePoolName := []string{ctx.Param("node_pool_name")}
	if err := o.bindNodePoolName(rNodePoolName, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteNodePoolParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ClusterID = raw

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteNodePoolParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Pattern("cluster_id", "path", o.ClusterID, `^[a-z][a-z0-9-:]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

func (o *DeleteNodePoolParams) bindNodePoolName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.NodePoolName = raw

	if err := o.validateNodePoolName(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteNodePoolParams) validateNodePoolName(formats strfmt.Registry) error {

	if err := validate.Pattern("node_pool_name", "path", o.NodePoolName, `^[a-z][a-z0-9-]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}