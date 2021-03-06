// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under 
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>. 
//
// https://documize.com

/*jshint node:true*/
module.exports = {
	name: 'intercom',

	contentFor: function (type, config) {
		// Not using Intercom for app activity logging?
		if (!config.APP.auditEnabled) {
			return '';
		}

		var appId = config.APP.intercomKey;

		if ('head-footer' === type && appId.length > 0) {
			return "<script type='text/javascript'> (function() { var w = window; var ic = w.Intercom; if (typeof ic === 'function') { ic('reattach_activator'); ic('update', intercomSettings); } else { var d = document; var i = function() { i.c(arguments) }; i.q = []; i.c = function(args) { i.q.push(args) }; w.Intercom = i; function l() { var s = d.createElement('script'); s.type = 'text/javascript'; s.async = true; s.src = 'https://widget.intercom.io/widget/${appId}'; var x = d.getElementsByTagName('script')[0]; x.parentNode.insertBefore(s, x); } if (w.attachEvent) { w.attachEvent('onload', l); } else { w.addEventListener('load', l, false); } } })() </script>";
		}
	},

	isDevelopingAddon: function () {
		return true;
	}
};