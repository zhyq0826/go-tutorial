import Route from '@ember/routing/route';
import Ember from 'ember';

export default Route.extend({
    model(params){
        return Ember.RSVP.hash({
            app: this.store.findOne('app', params.id),
            computer: this.store.find('computer', {appid: params.id})
        });
    }
});
