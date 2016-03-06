// Copyright 2016 Vincent Landgraf. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package confd

// ObjectMeta confd object metadata
type ObjectMeta struct {
	Ref      string `json:"ref"`
	Class    string `json:"class"`
	Type     string `json:"type"`
	Hidden   Bool   `json:"hidden"`
	Lock     string `json:"lock"`
	Nodel    string `json:"nodel"`
	Autoname Bool   `json:"autoname"`
}

// AnyObject a type that works with any confd object
type AnyObject struct {
	ObjectMeta
	Data map[string]interface{} `json:"data"`
}

// ChangeObject changes the object ref attributes
func (c *Conn) ChangeObject(ref string, attributes interface{}) (err error) {
	_, err = c.SimpleRequest("change_object", ref, attributes)
	return
}

// GetAnyObject returns a AnyObject for the given ref or nil
func (c *Conn) GetAnyObject(ref string) (*AnyObject, error) {
	response := new(AnyObject)
	err := c.Request("get_object", response, ref)
	return response, err
}
