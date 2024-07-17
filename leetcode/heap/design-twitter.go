package heap

import (
	"container/heap"
)

type Post struct {
	id      int
	tweetId int
}

type TweetId struct {
	post   Post
	userId int
	idx    int
}

type PostQueue []TweetId

func (pq PostQueue) Len() int           { return len(pq) }
func (pq PostQueue) Less(i, j int) bool { return pq[i].post.id > pq[j].post.id }
func (pq PostQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PostQueue) Push(x any) {
	tweet := x.(TweetId)
	*pq = append(*pq, tweet)
}
func (pq *PostQueue) Pop() any {
	old := *pq
	tweet := old[len(old)-1]
	old = old[:len(old)-1]
	*pq = old
	return tweet
}

type Twitter struct {
	count int
	users map[int]map[int]bool
	posts map[int][]Post
}

func Constructor() Twitter {
	return Twitter{
		count: 0,
		users: map[int]map[int]bool{},
		posts: map[int][]Post{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.posts[userId] = append(this.posts[userId], Post{
		id:      this.count,
		tweetId: tweetId,
	})
	this.count++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	maxHeaps := PostQueue{}
	for user := range this.users[userId] {
		if len(this.posts[user]) > 0 {
			index := len(this.posts[user]) - 1
			post := this.posts[user][index]
			maxHeaps = append(maxHeaps, TweetId{
				post:   post,
				idx:    index,
				userId: user,
			})
		}
	}
	// add the self post also
	if len(this.posts[userId]) > 0 {
		index := len(this.posts[userId]) - 1
		post := this.posts[userId][index]
		maxHeaps = append(maxHeaps, TweetId{
			post:   post,
			idx:    index,
			userId: userId,
		})
	}

	heap.Init(&maxHeaps)
	for len(maxHeaps) > 0 && len(res) < 10 {
		tweet := heap.Pop(&maxHeaps).(TweetId)
		res = append(res, tweet.post.tweetId)
		if tweet.idx > 0 {
			post := this.posts[tweet.userId][tweet.idx-1]
			heap.Push(&maxHeaps, TweetId{
				post:   post,
				userId: tweet.userId,
				idx:    tweet.idx - 1,
			})
		}
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if m, ok := this.users[followerId]; ok {
		m[followeeId] = true
		this.users[followerId] = m
	} else {
		m := map[int]bool{}
		m[followeeId] = true
		this.users[followerId] = m
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	m := this.users[followerId]
	if _, ok := m[followeeId]; ok {
		delete(m, followeeId)
		this.users[followerId] = m
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
