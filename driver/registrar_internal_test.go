package seqleasedriver

import (
	"testing"
)

func TestInternalRegistrar(t *testing.T) {

	var registrar Registrar = new(internalRegistrar)


	{
		const lessorCreatorName = "apple"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}




	{
		const lessorCreatorName = "apple"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}
	}




	{
		const lessorCreatorName = "apple"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		} else if nil != lessor {
			t.Errorf("Did not expect a lessor, but actually got one: (%T) %v", lessor, lessor)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		if lessor, err := registrar.Lessor(lessorCreatorName); nil != err {
			t.Errorf("Did not expect an error, but actually get one: (%T) %v", err, err)
			return
		} else if nil == lessor {
			t.Errorf("Did not expect a nil lessor, but actually got: (%T) %v", lessor)
			return
		}
	}



	{
		const lessorCreatorName = "apple"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "banana"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "cherry"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
	{
		const lessorCreatorName = "date"

		var lessorCreator LessorCreator = &MockLessorCreator{
			LessorFunc: func(...interface{}) (Lessor, error){
				return new(MockLessor), nil
			},
		}

		if err := registrar.Register(lessorCreatorName, lessorCreator); nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
	}
}
