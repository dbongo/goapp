describe('posts', function () {
  beforeEach(module('app'))
  var Posts, $httpBackend

  beforeEach(inject(function (_Posts_, _$httpBackend_) {
    Posts = _Posts_
    $httpBackend = _$httpBackend_
  }))

  afterEach(function (){
    $httpBackend.flush()
  })

  describe('#fetch', function () {
    beforeEach(function() {
      $httpBackend.expect('GET', '/api/posts')
      .respond([
        {username: 'dbongo', body: 'first post'},
        {username: 'dbongo', body: 'second post'}
      ])
    })

    it('gets 2 posts', function () {
      Posts.fetch()
      .then(function (posts) {
        expect(posts).to.have.length(2)
      })
    })
  })
})
