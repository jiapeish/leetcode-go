package _355_design_twitter

import (
	"sort"
	"time"
)

type tweet struct {
	id int
	time int64
}

type tweets []tweet

func (t tweets) Len() int {
	return len(t)
}

func (t tweets) Less(i, j int) bool {
	return t[i].time > t[j].time
}

func (t tweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Twitter represents twitter user
type Twitter struct {
	userTweets map[int]tweets
	following  map[int][]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	t := make(map[int]tweets, 1)
	f := make(map[int][]int, 1)

	return Twitter{
		userTweets: t,
		following:  f,
	}
}


/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int)  {
	t.userTweets[userId] = append(t.userTweets[userId], tweet{
		id:   tweetId,
		time: time.Now().UnixNano(),
	})
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	// get own tweets
	ts := make(tweets, 0, len(t.userTweets[userId]))
	ts = append(ts, t.userTweets[userId]...)

	// get followee' tweets
	for _, id := range t.following[userId] {
		ts = append(ts, t.userTweets[id]...)
	}

	sort.Sort(ts)

	res := make([]int, 0, 10)
	for i := 0; i < len(ts) && i < 10; i++ {
		res = append(res, ts[i].id)
	}

	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int)  {
	if followeeId == followerId {
		return
	}

	for _, id := range t.following[followerId] {
		if followeeId == id {
			return
		}
	}
	t.following[followerId] = append(t.following[followerId], followeeId)
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int)  {
	for i, id := range t.following[followerId] {
		if followeeId == id {
			t.following[followerId] = append(t.following[followerId][:i], t.following[followerId][i+1:]...)
		}
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