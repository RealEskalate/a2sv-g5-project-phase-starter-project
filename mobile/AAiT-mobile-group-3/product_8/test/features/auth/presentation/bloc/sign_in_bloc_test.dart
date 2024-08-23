// import 'package:bloc_test/bloc_test.dart';
// import 'package:dartz/dartz.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/failure/failure.dart';
// import 'package:product_8/features/auth/domain/entities/sign_in_user_entitiy.dart';
// import 'package:product_8/features/auth/domain/entities/user_data_entity.dart';
// import 'package:product_8/features/auth/presentation/bloc/sign_in_bloc/sign_in_bloc.dart';
// import 'package:product_8/features/auth/presentation/bloc/sign_in_bloc/sign_in_event.dart';
// import 'package:product_8/features/auth/presentation/bloc/sign_in_bloc/sign_in_state.dart';

// import '../../../../helpers/test_helper.mocks.dart';


// void main() {
//   late MockSignInUseCase mockSignInUseCase;
//   late SignInBloc signInBloc;

//   const SignInUserEntitiy testUser = SignInUserEntitiy(
//     email: 'test@example.com',
//     password: 'password123',
//   );

//   const UserDataEntity testUserData = UserDataEntity(
//     data: Data(name: 'Test User', email: 'test@example.com'),
//     token: '123token',
//   );

//   setUp(() {
//     mockSignInUseCase = MockSignInUseCase();
//     signInBloc = SignInBloc(signInUseCase: mockSignInUseCase);
//   });

//   test('initial state should be SignInInitial', () {
//     expect(signInBloc.state, equals(SignInInitial()));
//   });

//   group('SignIn Event', () {
//     blocTest<SignInBloc, SignInState>(
//       'emits [SignInLoadingState, SignInLoadedState] when SignIn is successful',
//       build: () {
//         when(mockSignInUseCase.call(any))
//             .thenAnswer((_) async => const Right(testUserData));
//         return signInBloc;
//       },
//       act: (bloc) => bloc.add(const OnLogInButtonPressedEvent(signInUserEntitiy: testUser)),
//       wait: const Duration(milliseconds: 500),
//       expect: () => [
//         SignInLoadingState(),
//         const SignInLoadedState(userDataEntity: testUserData),
//       ],
//       verify: (_) {
//         verify(mockSignInUseCase.call(any)).called(1);
//       },
//     );

//     blocTest<SignInBloc, SignInState>(
//       'emits [SignInLoadingState, SignInErrorState] when SignIn fails due to ServerFailure',
//       build: () {
//         when(mockSignInUseCase.call(any))
//             .thenAnswer((_) async => const Left(ServerFailure(message: 'Server failure')));
//         return signInBloc;
//       },
//       act: (bloc) => bloc.add(const OnLogInButtonPressedEvent(signInUserEntitiy:  testUser)),
//       wait: const Duration(milliseconds: 500),
//       expect: () => [
//         SignInLoadingState(),
//         const SignInErrorState(message: SERVER_FAILURE_MESSAGE),
//       ],
//       verify: (_) {
//         verify(mockSignInUseCase.call(any)).called(1);
//       },
//     );

//     blocTest<SignInBloc, SignInState>(
//       'emits [SignInLoadingState, SignInErrorState] when SignIn fails due to ConnectionFailure',
//       build: () {
//         when(mockSignInUseCase.call(any))
//             .thenAnswer((_) async => const Left(ConnectionFailure(message: 'Connection failure')));
//         return signInBloc;
//       },
//       act: (bloc) => bloc.add(const OnLogInButtonPressedEvent(signInUserEntitiy:  testUser)),
//       wait: const Duration(milliseconds: 500),
//       expect: () => [
//         SignInLoadingState(),
//         const SignInErrorState(message: CONNECTION_FAILURE_MESSAGE),
//       ],
//       verify: (_) {
//         verify(mockSignInUseCase.call(any)).called(1);
//       },
//     );
//   });
// }
