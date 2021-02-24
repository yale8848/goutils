package inject

var defaultInjecter Graph

func Inject(vs ...interface{}) error {
	for _, v := range vs {
		err := defaultInjecter.Provide(&Object{Value: v})
		if err != nil {
			return err
		}
	}
	return defaultInjecter.Populate()
}
func InjectByName(name string, v interface{}) error {
	err := defaultInjecter.Provide(&Object{Value: v, Name: name})
	if err != nil {
		return err
	}
	return defaultInjecter.Populate()
}
