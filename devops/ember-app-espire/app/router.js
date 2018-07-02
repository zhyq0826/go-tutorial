import EmberRouter from '@ember/routing/router';
import config from './config/environment';

const Router = EmberRouter.extend({
  location: config.locationType,
  rootURL: config.rootURL
});

Router.map(function() {
  this.route('domain');
  this.route('computer');
  this.route('app', function() {
    this.route('detail',{path: '/:id/detail'}, function() {
      this.route('tags');
    });
  });
});

export default Router;
