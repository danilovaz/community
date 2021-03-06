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

import Ember from 'ember';
import TooltipMixin from '../../mixins/tooltip';

const {
	computed
} = Ember;

export default Ember.Component.extend(TooltipMixin, {
	deleteChildren: false,

	checkId: computed('page', function () {
		let id = this.get('page.id');
		return `delete-check-button-${id}`;
	}),

	dropTarget: computed('page', function () {
		let id = this.get('page.id');
		return `delete-page-button-${id}`;
	}),

	didRender() {
		if (this.get('isEditor')) {
			let self = this;
			$(".page-edit-button, .page-delete-button").each(function (i, el) {
				self.addTooltip(el);
			});
		}
	},

	willDestroyElement() {
		this.destroyTooltips();
	},

	actions: {
		editPage(id) {
			this.attrs.onEditPage(id);
		},

		deletePage(id) {
			this.attrs.onDeletePage(id, this.get('deleteChildren'));
		},
	}
});
