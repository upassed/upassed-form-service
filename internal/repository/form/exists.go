package form

import "context"

func (repository *formRepositoryImpl) ExistsByNameAndTeacherUsername(ctx context.Context, formName, teacherUsername string) (bool, error) {
	return false, nil
}