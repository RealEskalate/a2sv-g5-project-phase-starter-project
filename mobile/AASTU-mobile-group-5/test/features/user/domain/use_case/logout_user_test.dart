import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/features/user/domain/use_case/logout_user.dart';
import '../../../product/data/datasources/product_remote_data_source_test.mocks.dart';

void main() {
  late LogOut logOut;
  late MockUserLocalDataSource mockUserLocalDataSource;

  setUp(() {
    mockUserLocalDataSource = MockUserLocalDataSource();
    logOut = LogOut(mockUserLocalDataSource);
  });

  test('should call deleteAccessToken on the local data source', () async {
    // Arrange
    when(mockUserLocalDataSource.deleteAccessToken()).thenAnswer((_) async => Future.value());

    // Act
    await logOut();

    // Assert
    verify(mockUserLocalDataSource.deleteAccessToken());
    verifyNoMoreInteractions(mockUserLocalDataSource);
  });
}
