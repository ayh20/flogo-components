package end

// Input data structure
type Input struct {
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	return nil
}

//Output data structure
type Output struct {
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}

//FromMap Output  from map
func (o *Output) FromMap(values map[string]interface{}) error {

	return nil
}
