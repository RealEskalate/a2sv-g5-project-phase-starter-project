import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';
import 'package:ecom_app/features/auth/domain/entities/login_entity.dart';
import 'package:ecom_app/features/auth/domain/entities/register_entity.dart';
import 'package:ecom_app/features/auth/domain/entities/user_data_entity.dart';
import 'package:ecom_app/features/auth/domain/usecases/login.dart';
import 'package:ecom_app/features/auth/domain/usecases/register.dart';
import 'package:ecom_app/features/auth/presentation/bloc/auth_bloc.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../helpers/test_helper.mocks.dart';

void main() {
  late MockLoginUsecase mockLoginUsecase;
  late MockRegisterUsecase mockRegisterUsecase;
  late MockLogoutUsecase mockLogoutUsecase;
  late MockGetUserUsecase mockGetUserUsecase;
  late AuthBloc authBloc;

  setUp(() {
    mockLoginUsecase = MockLoginUsecase();
    mockRegisterUsecase = MockRegisterUsecase();
    mockLogoutUsecase = MockLogoutUsecase();
    mockGetUserUsecase = MockGetUserUsecase();
    authBloc = AuthBloc(
      mockLoginUsecase,
      mockRegisterUsecase,
      mockGetUserUsecase,
      mockLogoutUsecase,
    );
  });

  final tLoginEntity = LoginEntity(email: 'email', password: 'password');
  final tRegistrationEntity =
      RegistrationEntity(email: 'email', password: 'password', name: 'name');
  final tUserDataEntity = UserDataEntity(email: 'email', name: 'name');
  test('initial state should be AuthInitial', () {
    expect(authBloc.state, AuthInitial());
  });

  group('LoginEvent', () {
    blocTest<AuthBloc, AuthState>('emits [AuthLoading, AuthSuccess]',
        build: () {
          return authBloc;
        },
        setUp: () {
          when(mockLoginUsecase(LoginParams(loginEntity: tLoginEntity)))
              .thenAnswer((_) async => const Right(unit));
          return authBloc;
        },
        act: (bloc) => bloc.add(LoginEvent(loginEntity: tLoginEntity)),
        expect: () => [AuthLoading(), AuthSuccess()]);

    blocTest(
      'emits [AuthLoading, AuthError]',
      build: () {
        return authBloc;
      },
      setUp: () {
        when(mockLoginUsecase(LoginParams(loginEntity: tLoginEntity)))
            .thenAnswer((_) async => const Left(ServerFailure('error')));
        return authBloc;
      },
      act: (bloc) => bloc.add(LoginEvent(loginEntity: tLoginEntity)),
      expect: () => [AuthLoading(), AuthError(message: 'error')],
    );
  });
  group('RegisterEvent', () {
    blocTest<AuthBloc, AuthState>('emits [AuthLoading, AuthSuccess]',
        build: () {
          return authBloc;
        },
        setUp: () {
          when(mockRegisterUsecase(
                  RegisterParams(registrationEntity: tRegistrationEntity)))
              .thenAnswer((_) async => const Right(unit));
          return authBloc;
        },
        act: (bloc) =>
            bloc.add(RegisterEvent(registrationEntity: tRegistrationEntity)),
        expect: () => [AuthLoading(), AuthRegisterSuccess()]);

    blocTest(
      'emits [AuthLoading, AuthError]',
      build: () {
        return authBloc;
      },
      setUp: () {
        when(mockRegisterUsecase(
                RegisterParams(registrationEntity: tRegistrationEntity)))
            .thenAnswer((_) async => const Left(ServerFailure('error')));
        return authBloc;
      },
      act: (bloc) =>
          bloc.add(RegisterEvent(registrationEntity: tRegistrationEntity)),
      expect: () => [AuthLoading(), AuthError(message: 'error')],
    );
  });

  group('GetUserEvent', () {
    blocTest<AuthBloc, AuthState>('emits [AuthLoading, AuthGetUserSuccess]',
        build: () {
          return authBloc;
        },
        setUp: () {
          when(mockGetUserUsecase(NoParams()))
              .thenAnswer((_) async => Right(tUserDataEntity));
          return authBloc;
        },
        act: (bloc) => bloc.add(GetUserEvent()),
        expect: () => [
              AuthLoading(),
              AuthAuthenticated(userDataEntity: tUserDataEntity)
            ]);

    blocTest(
      'emits [AuthLoading, AuthError]',
      build: () {
        return authBloc;
      },
      setUp: () {
        when(mockGetUserUsecase(NoParams()))
            .thenAnswer((_) async => const Left(ServerFailure('error')));
        return authBloc;
      },
      act: (bloc) => bloc.add(GetUserEvent()),
      expect: () => [AuthLoading(), AuthError(message: 'error')],
    );
  });

  group('LogoutEvent', () {
    blocTest<AuthBloc, AuthState>('emits [AuthLoading, AuthLoggedOut]',
        build: () {
          return authBloc;
        },
        setUp: () {
          when(mockLogoutUsecase(NoParams()))
              .thenAnswer((_) async => const Right(unit));
          return authBloc;
        },
        act: (bloc) => bloc.add(LogoutEvent()),
        expect: () => [AuthLoading(), AuthLoggedOut()]);

    blocTest(
      'emits [AuthLoading, AuthError]',
      build: () {
        return authBloc;
      },
      setUp: () {
        when(mockLogoutUsecase(NoParams()))
            .thenAnswer((_) async => const Left(ServerFailure('error')));
        return authBloc;
      },
      act: (bloc) => bloc.add(LogoutEvent()),
      expect: () => [AuthLoading(), AuthError(message: 'error')],
    );
  });
}
