//nolint:revive
package tier0s

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra/tier_0s"
	model1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_0s"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

type FloodProtectionProfileBindingMapClientContext utl.ClientContext

func NewFloodProtectionProfileBindingsClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *FloodProtectionProfileBindingMapClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewFloodProtectionProfileBindingsClient(connector)

	case utl.Global:
		client = client1.NewFloodProtectionProfileBindingsClient(connector)

	default:
		return nil
	}
	return &FloodProtectionProfileBindingMapClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID}
}

func (c FloodProtectionProfileBindingMapClientContext) Get(tier0IdParam string, floodProtectionProfileBindingIdParam string) (model0.FloodProtectionProfileBindingMap, error) {
	var obj model0.FloodProtectionProfileBindingMap
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.FloodProtectionProfileBindingsClient)
		obj, err = client.Get(tier0IdParam, floodProtectionProfileBindingIdParam)
		if err != nil {
			return obj, err
		}

	case utl.Global:
		client := c.Client.(client1.FloodProtectionProfileBindingsClient)
		gmObj, err1 := client.Get(tier0IdParam, floodProtectionProfileBindingIdParam)
		if err1 != nil {
			return obj, err1
		}
		var rawObj interface{}
		rawObj, err = utl.ConvertModelBindingType(gmObj, model1.FloodProtectionProfileBindingMapBindingType(), model0.FloodProtectionProfileBindingMapBindingType())
		obj = rawObj.(model0.FloodProtectionProfileBindingMap)

	default:
		return obj, errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c FloodProtectionProfileBindingMapClientContext) Delete(tier0IdParam string, floodProtectionProfileBindingIdParam string) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.FloodProtectionProfileBindingsClient)
		err = client.Delete(tier0IdParam, floodProtectionProfileBindingIdParam)

	case utl.Global:
		client := c.Client.(client1.FloodProtectionProfileBindingsClient)
		err = client.Delete(tier0IdParam, floodProtectionProfileBindingIdParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c FloodProtectionProfileBindingMapClientContext) Patch(tier0IdParam string, floodProtectionProfileBindingIdParam string, floodProtectionProfileBindingMapParam model0.FloodProtectionProfileBindingMap) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.FloodProtectionProfileBindingsClient)
		err = client.Patch(tier0IdParam, floodProtectionProfileBindingIdParam, floodProtectionProfileBindingMapParam)

	case utl.Global:
		client := c.Client.(client1.FloodProtectionProfileBindingsClient)
		gmObj, err1 := utl.ConvertModelBindingType(floodProtectionProfileBindingMapParam, model0.FloodProtectionProfileBindingMapBindingType(), model1.FloodProtectionProfileBindingMapBindingType())
		if err1 != nil {
			return err1
		}
		err = client.Patch(tier0IdParam, floodProtectionProfileBindingIdParam, gmObj.(model1.FloodProtectionProfileBindingMap))

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c FloodProtectionProfileBindingMapClientContext) Update(tier0IdParam string, floodProtectionProfileBindingIdParam string, floodProtectionProfileBindingMapParam model0.FloodProtectionProfileBindingMap) (model0.FloodProtectionProfileBindingMap, error) {
	var err error
	var obj model0.FloodProtectionProfileBindingMap

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.FloodProtectionProfileBindingsClient)
		obj, err = client.Update(tier0IdParam, floodProtectionProfileBindingIdParam, floodProtectionProfileBindingMapParam)

	case utl.Global:
		client := c.Client.(client1.FloodProtectionProfileBindingsClient)
		gmObj, err := utl.ConvertModelBindingType(floodProtectionProfileBindingMapParam, model0.FloodProtectionProfileBindingMapBindingType(), model1.FloodProtectionProfileBindingMapBindingType())
		if err != nil {
			return obj, err
		}
		gmObj, err = client.Update(tier0IdParam, floodProtectionProfileBindingIdParam, gmObj.(model1.FloodProtectionProfileBindingMap))
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.FloodProtectionProfileBindingMapBindingType(), model0.FloodProtectionProfileBindingMapBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.FloodProtectionProfileBindingMap)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}