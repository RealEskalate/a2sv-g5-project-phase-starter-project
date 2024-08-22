// import 'package:bloc_test/bloc_test.dart';
// import 'package:dartz/dartz.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/failure/failure.dart';
// import 'package:product_8/features/auth/domain/entities/sign_up_user_entitiy.dart';
// import 'package:product_8/features/auth/presentation/bloc/sign_up_bloc/sign_up_bloc.dart';
// import 'package:product_8/features/auth/presentation/bloc/sign_up_bloc/sign_up_event.dart';
// import 'package:product_8/features/auth/presentation/bloc/sign_up_bloc/sign_up_state.dart';

// import '../../../../helpers/test_helper.mocks.dart';



// void main() {
//   late MockSignUpUseCase mockSignUpUseCase;
//   late SignUpBloc signUpBloc;
//  late MockSignInUseCase mockSignInUseCase;

//   const SignUpUserEntitiy testUser = SignUpUserEntitiy(
//     email: 'test@example.com',
//     password: 'password123',
//     name: 'Test User',
//   );

//   const SignUpUserEntitiy testUserData = SignUpUserEntitiy(
//     email: 'test@example.com',
//     password: 'password123',
//     name: 'Test User',
//   );

//   // const UserDataEntity testUserData = UserDataEntity(
//   //   data: Data(name: 'Test User', email: 'test@example.com'),
//   //   token: '123token',
//   // );

//   setUp(() {
//     mockSignUpUseCase = MockSignUpUseCase();
//     mockSignInUseCase = MockSignInUseCase();
//     signUpBloc = SignUpBloc(signInUseCase: mockSignInUseCase, signUpUseCase: mockSignUpUseCase);  // Null for signInUseCase since it's not used here
//   });

//   test('initial state should be SignUpInitialState', () {
//     expect(signUpBloc.state, equals(SignUpInitialState()));
//   });

//   group('SignUp Event', () {
//     blocTest<SignUpBloc, SignUpState>(
//       'emits [SignUpLoadingState, SignUpLoadedState] when SignUp is successful',
//       build: () {
//         when(mockSignUpUseCase.call(any))
//             .thenAnswer((_) async => const Right(testUserData));
//         return signUpBloc;
//       },
//       act: (bloc) => bloc.add(const OnSignUpButtonPressedEvent(signUpUserEntitiy:  testUser)),
//       wait: const Duration(milliseconds: 500),
//       expect: () => [
//         SignUpLoadingState(),
//         const SignUpLoadedState(signUpUserEnity: testUserData),
//       ],
//       verify: (_) {
//         verify(mockSignUpUseCase.call(any)).called(1);
//       },
//     );

//     blocTest<SignUpBloc, SignUpState>(
//       'emits [SignUpLoadingState, SignUpErrorState] when SignUp fails due to ServerFailure',
//       build: () {
//         when(mockSignUpUseCase.call(any))
//             .thenAnswer((_) async => const Left(ServerFailure(message: 'Server failure')));
//         return signUpBloc;
//       },
//       act: (bloc) => bloc.add(const OnSignUpButtonPressedEvent(signUpUserEntitiy:  testUser)),
//       wait: const Duration(milliseconds: 500),
//       expect: () => [
//         SignUpLoadingState(),
//         const SignUpErrorState(message: SERVER_FAILURE_MESSAGE),
//       ],
//       verify: (_) {
//         verify(mockSignUpUseCase.call(any)).called(1);
//       },
//     );

//     blocTest<SignUpBloc, SignUpState>(
//       'emits [SignUpLoadingState, SignUpErrorState] when SignUp fails due to ConnectionFailure',
//       build: () {
//         when(mockSignUpUseCase.call(any))
//             .thenAnswer((_) async => const Left(ConnectionFailure(message: 'Connection failure')));
//         return signUpBloc;
//       },
//       act: (bloc) => bloc.add(const OnSignUpButtonPressedEvent(signUpUserEntitiy:  testUser)),
//       wait: const Duration(milliseconds: 500),
//       expect: () => [
//         SignUpLoadingState(),
//         const SignUpErrorState(message: CONNECTION_FAILURE_MESSAGE),
//       ],
//       verify: (_) {
//         verify(mockSignUpUseCase.call(any)).called(1);
//       },
//     );
//   });
// }
