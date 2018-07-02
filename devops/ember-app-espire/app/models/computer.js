import EmberObject from '@ember/object';
// import A from '@ember/array';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;


export default EmberObject.extend(model, {
    url: "/v1/computer",
    urlForFind: function() {
        return this.get('api') + '/list';
    },
    init(){
        this._super(...arguments);
        this.model = {
            'name': attr('string')
        }
    }
});