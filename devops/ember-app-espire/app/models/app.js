import EmberObject from '@ember/object';
// import A from '@ember/array';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;


export default EmberObject.extend(model, {
    url: "/v1/app",
    urlForFind: function() {
        return this.get('api') + '/list';
    },
    init(){
        this._super(...arguments);
        this.model = {
            name: attr('string'),
            url: attr('string'),
            desc: attr('string'),
            repository_url: attr('string'),
            deploy_dir: attr('string'),
            monitor_url: attr('string'),
        }
    }
});