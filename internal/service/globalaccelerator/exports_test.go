// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalaccelerator

// Exports for use in tests only.
var (
	ResourceAccelerator              = resourceAccelerator
	ResourceCrossAccountAttachment   = newCrossAccountAttachmentResource
	ResourceCustomRoutingAccelerator = resourceCustomRoutingAccelerator
	ResourceCustomRoutingListener    = resourceCustomRoutingListener
	ResourceEndpointGroup            = resourceEndpointGroup
	ResourceListener                 = resourceListener

	FindAcceleratorByARN              = findAcceleratorByARN
	FindCrossAccountAttachmentByARN   = findCrossAccountAttachmentByARN
	FindCustomRoutingAcceleratorByARN = findCustomRoutingAcceleratorByARN
	FindCustomRoutingListenerByARN    = findCustomRoutingListenerByARN
	FindListenerByARN                 = findListenerByARN
)
