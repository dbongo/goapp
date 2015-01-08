angular.module('hackapp').controller('PostsCtrl', function PostsCtrl(Posts, Auth) {
  var vm = this

  vm.posts = []
  vm.addPost = addPost

  Posts.fetch().then(function(posts) {
    vm.posts = posts
  })

  function addPost() {
    if (vm.postBody) {
      Posts.create({
        username: Auth.getCurrentUser().username,
        body: vm.postBody
      })
      .then(function(res) {
        vm.postBody = null
        vm.posts.push(res.data)
      })
    }
  }
})
