
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__Name struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_interface  struct {
			XMLName xml.Name `xml:"interface"`
			V_name  *string  `xml:"name,omitempty"`
			V_unit  struct {
				XMLName xml.Name `xml:"unit"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_storm__control  struct {
					XMLName xml.Name `xml:"storm-control"`
					V_profile__name  *string  `xml:"profile-name,omitempty"`
				} `xml:"family>ethernet-switching>storm-control"`
			} `xml:"unit"`
		} `xml:"interfaces>interface"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_profile__name := d.Get("profile__name").(string)


	config := xmlInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_unit.V_name__1 = &V_name__1
	config.Groups.V_interface.V_unit.V_storm__control.V_profile__name = &V_profile__name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameRead(d,m)
}

func junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__Name{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_interface.V_name)
	d.Set("name__1", config.Groups.V_interface.V_unit.V_name__1)
	d.Set("profile__name", config.Groups.V_interface.V_unit.V_storm__control.V_profile__name)

	return nil
}

func junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_profile__name := d.Get("profile__name").(string)


	config := xmlInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_unit.V_name__1 = &V_name__1
	config.Groups.V_interface.V_unit.V_storm__control.V_profile__name = &V_profile__name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameRead(d,m)
}

func junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__Name() *schema.Resource {
	return &schema.Resource{
		Create: junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameCreate,
		Read: junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameRead,
		Update: junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameUpdate,
		Delete: junosInterfacesInterfaceUnitFamilyEthernet__SwitchingStorm__ControlProfile__NameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface.V_unit",
			},
			"profile__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface.V_unit.V_storm__control. Profile name",
			},
		},
	}
}