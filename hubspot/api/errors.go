package hubspot

import "errors"

var ErrBatchCreate = errors.New("error creating contacts")

var ErrPropertyValidation = errors.New("error validating properties")

var ErrObjectAlreadyExists = errors.New("object already exists")

var ErrApiCall = errors.New("error making api call")

var ErrAlreadyExists = errors.New("resource already exists")

var ErrResourceAlreadyExists = errors.New("resource already exists")

var ErrSubscriptionAlreadyUnsubscribed = errors.New("subscription already unsubscribed")

var ErrSubscriptionAlreadySubscribed = errors.New("subscription already subscribed")
