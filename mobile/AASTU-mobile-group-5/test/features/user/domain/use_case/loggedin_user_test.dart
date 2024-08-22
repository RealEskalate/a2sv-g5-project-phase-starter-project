import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/features/user/domain/use_case/loggedin_user.dart';
import '../../../product/data/datasources/product_remote_data_source_test.mocks.dart';

void main() {
  late IsLoggedIn isLoggedIn;
  late MockUserLocalDataSource mockUserLocalDataSource;

  setUp(() {
    mockUserLocalDataSource = MockUserLocalDataSource();
    isLoggedIn = IsLoggedIn(mockUserLocalDataSource);
  });

  test('should return true when there is a valid token', () async {
    // Arrange
    when(mockUserLocalDataSource.getAccessToken()).thenAnswer((_) async => 'valid_token');

    // Act
    final result = await isLoggedIn();

    // Assert
    expect(result, true);
  });

  test('should return false when there is no token', () async {
    // Arrange
    when(mockUserLocalDataSource.getAccessToken()).thenAnswer((_) async => null);

    // Act
    final result = await isLoggedIn();

    // Assert
    expect(result, false);
  });
}
