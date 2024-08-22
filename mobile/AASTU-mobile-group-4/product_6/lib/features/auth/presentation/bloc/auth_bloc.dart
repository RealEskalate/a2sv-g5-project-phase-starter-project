import 'dart:async';

import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/no_param_use_cases.dart';
import '../../domain/entity/auth_entity.dart';
import '../../domain/usecases/get_user_profile.dart';
import '../../domain/usecases/login_usecase.dart';
import '../../domain/usecases/logout_usecase.dart';
import '../../domain/usecases/register_usecase.dart';
import '../widgets/logo.dart';
import 'auth_event.dart';
import 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final Login loginUseCase;
  final Register registerUseCase;
  final GetUserProfile getUserProfileUseCase;
  final LogoutUseCase logoutUseCase;

  AuthBloc({
    required this.loginUseCase,
    required this.registerUseCase,
    required this.getUserProfileUseCase,
    required this.logoutUseCase,
  }) : super(AuthInitial()) {
    on<LoginEvent>(_onLoginEvent);
    on<RegisterEvent>(_onRegisterEvent);
    on<GetUserProfileEvent>(_onGetUserProfileEvent);
    on<LogoutEvent>(_onLogoutEvent);
  }

  void _onLoginEvent(LoginEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());

    final result = await loginUseCase(AuthEntity(
      email: event.email,
      password: event.password,
    ));

    result.fold(
      (failure) => emit(AuthError(message: _mapFailureToMessage(failure))),
      (authResponse) => emit(AuthLoaded(authResponse: authResponse)),
    );
  }

  void _onRegisterEvent(RegisterEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    print(event.senduserentity);
    final result = await registerUseCase(
      event.senduserentity,
    );
    print('from auth_bloc');
    print(result);

    result.fold(
      (failure) => emit(AuthError(message: _mapFailureToMessage(failure))),
      (user) => emit(UserSuccessRegister()),
    );
  }

  void _onGetUserProfileEvent(
      GetUserProfileEvent event, Emitter<AuthState> emit) async {
    emit(AuthLoading());

    final result = await getUserProfileUseCase(event);
    print('from auth_bloc');
    print(result);

    result.fold(
      (failure) => emit(AuthError(message: _mapFailureToMessage(failure))),
      (user) => emit(UserProfileLoaded(user: user)),
    );
  }
  FutureOr<void> _onLogoutEvent(LogoutEvent event, Emitter<AuthState> emit) async{
    emit(AuthLoading());

    final result = await logoutUseCase(event);

    print('from auth_bloc');
    print(result);
    result.fold(
      (failure) => emit(AuthError(message: _mapFailureToMessage(failure))),
      (user) => emit(LogoutSuccess()),
    );

  }

  String _mapFailureToMessage(Failure failure) {
    switch (failure.runtimeType) {
      case ServerFailure:
        return 'Server Failure';
      case ServerFailure(message: 'Invalid Credentials'):
        return 'Cache Failure';
      default:
        return 'Unexpected Error';
    }
  }


  }

