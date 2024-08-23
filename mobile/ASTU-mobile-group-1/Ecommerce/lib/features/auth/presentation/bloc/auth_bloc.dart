import 'dart:developer';

import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/usecases/check_signed_in_usecase.dart';
import '../../domain/usecases/get_user.dart';
import '../../domain/usecases/log_out_usecase.dart';
import '../../domain/usecases/sign_in_usecase.dart';
import '../../domain/usecases/sign_up_usecase.dart';
import 'auth_event.dart';
import 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final SignInUsecase _signInUsecase;
  final SignUpUsecase _signUpUsecase;
  final LogOutUsecase _logOutUsecase;
  final CheckSignedInUsecase _checkSignedInUsecase;
  final GetUserUsecase _getUserUsecase;

  AuthBloc({
    required SignInUsecase signInUsecase,
    required SignUpUsecase signUpUsecase,
    required LogOutUsecase logOutUsecase,
    required CheckSignedInUsecase checkSignedInUsecase,
    required GetUserUsecase getUserUsecase,
  })  : _signInUsecase = signInUsecase,
        _signUpUsecase = signUpUsecase,
        _logOutUsecase = logOutUsecase,
        _checkSignedInUsecase = checkSignedInUsecase,
        _getUserUsecase = getUserUsecase,
        super(InitalState()) {
    on<SignInEvent>(_onSignInEvent);
    on<SignUpEvent>(_onSignUpEvent);
    on<LogOutEvent>(_onLogOut);
    on<CheckSignedInEvent>(_onCheckSignedIn);
    on<GetUserEvent>(_onGetUser);
  }

  Future<void> _onGetUser(GetUserEvent event, Emitter emit) async {
    emit(LoadingState());

    final result = await _getUserUsecase();
    result.fold((failure) {
      emit(ErrorState(message: failure.message));
      add(LogOutEvent());
    }, (data) {
      emit(UserIsReady());
    });
  }

  Future<void> _onSignInEvent(SignInEvent event, Emitter emit) async {
    emit(LoadingState());

    final result =
        await _signInUsecase(email: event.email, password: event.password);

    result.fold((failure) {
      emit(ErrorState(message: failure.message));
      // add();
    }, (data) {
      emit(SignedInState());
    });
  }

  Future<void> _onSignUpEvent(SignUpEvent event, Emitter emit) async {
    emit(LoadingState());

    final result = await _signUpUsecase(
        email: event.email,
        password: event.password,
        name: event.name,
        repeatedPassword: event.repeatedPassword);

    result.fold((failure) {
      emit(ErrorState(message: failure.message));
    }, (data) {
      emit(SignedUpState());
    });
  }

  Future<void> _onLogOut(LogOutEvent event, Emitter emit) async {
    emit(LoadingState());

    final result = await _logOutUsecase();

    result.fold((failure) {
      emit(
        ErrorState(message: failure.message),
      );
    }, (data) {
      log('Logged out!');
      emit(SignInState());
    });
  }

  Future<void> _onCheckSignedIn(CheckSignedInEvent event, Emitter emit) async {
    emit(LoadingState());

    final result = await _checkSignedInUsecase();

    result.fold((failure) {
      log(failure.message);
      emit(SignInState());
    }, (data) {
      if (data) {
        emit(SignedInState());
      } else {
        emit(SignInState());
      }
    });
  }
}
