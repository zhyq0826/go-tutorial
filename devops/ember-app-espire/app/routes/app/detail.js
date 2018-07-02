import Route from '@ember/routing/route';
import { hash } from 'rsvp';

export default Route.extend({
    model(params){
        return hash({
            app: this.store.findOne('app', params.id),
            computer: this.store.find('computer', {appid: params.id})
        });
    }
});
