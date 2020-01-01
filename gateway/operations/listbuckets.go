package operations

import (
	"net/http"
	"treeverse-lake/gateway/errors"
	"treeverse-lake/gateway/permissions"
	"treeverse-lake/gateway/serde"
)

type ListBuckets struct{}

func (controller *ListBuckets) GetArn() string {
	return "arn:treeverse:repos:::*"
}

func (controller *ListBuckets) GetPermission() string {
	return permissions.PermissionReadRepo
}

func (controller *ListBuckets) Handle(o *AuthenticatedOperation) {
	repos, err := o.Index.ListRepos()
	if err != nil {
		o.EncodeError(errors.Codes.ToAPIErr(errors.ErrInternalError))
		return
	}

	// assemble response
	buckets := make([]serde.Bucket, len(repos))
	for i, repo := range repos {
		buckets[i] = serde.Bucket{
			CreationDate: serde.Timestamp(repo.GetCreationDate()),
			Name:         repo.GetRepoId(),
		}
	}

	// write response
	o.EncodeResponse(serde.ListAllMyBucketsResult{
		Buckets: serde.Buckets{Bucket: buckets},
	}, http.StatusOK)

}
