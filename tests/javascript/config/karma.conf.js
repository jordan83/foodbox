basePath = '../';

files = [
  JASMINE,
  JASMINE_ADAPTER,
  '../../public/javascripts/angular.js',
  '../../public/javascripts/angular-*.js',
  'lib/angular-mocks.js',
  '../../public/javascripts/app/*.js',
  'unit/*.js'
];

autoWatch = true;

browsers = ['Chrome'];

junitReporter = {
  outputFile: 'test_out/unit.xml',
  suite: 'unit'
};