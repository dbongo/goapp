describe('posts.ctrl', function () {
  beforeEach(module('app'))
  var vm = this
  var mockPosts = {}

  beforeEach(inject(function ($q) {
    mockPosts.fetch = function () {
      var deferred = $q.defer()
      deferred.resolve([
        {username: 'dbongo', body: 'first post'},
        {username: 'dbongo', body: 'second post'}
      ])
      return deferred.promise
    }
    mockPosts.create = function () {
      var deferred = $q.defer()
      deferred.resolve()
      return deferred.promise
    }
  }))

  beforeEach(inject(function ($rootScope, $controller) {
    //$scope = $rootScope.$new()
    $controller('PostsCtrl', {
      //vm: $scope,
      Posts: mockPosts
    })
  }))

  it('loads posts from the service', function () {
    vm.$digest()
    expect(vm.posts).to.have.length(2)
  })

  it('sends a new post to the service', function () {
    sinon.spy(mockPosts, 'create')
    vm.postBody = 'my new post'
    vm.addPost()
    expect(mockPosts.create).to.have.been.calledWith({body: 'my new post'})
  })
})
