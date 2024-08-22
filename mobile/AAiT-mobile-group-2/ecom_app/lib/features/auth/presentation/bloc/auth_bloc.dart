import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/usecase/usecase.dart';
import '../../domain/entities/login_entity.dart';
import '../../domain/entities/register_entity.dart';
import '../../domain/entities/user_data_entity.dart';
import '../../domain/usecases/get_user.dart';
import '../../domain/usecases/login.dart';
import '../../domain/usecases/logout.dart';
import '../../domain/usecases/register.dart';

part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final LoginUsecase _loginUsecase;
  final RegisterUsecase _registerUsecase;
  final GetUserUsecase _getUserUsecase;
  final LogoutUsecase _logoutUsecase;

  AuthBloc(this._loginUsecase, this._registerUsecase, this._getUserUsecase,
      this._logoutUsecase)
      : super(AuthInitial()) {
    on<LoginEvent>((event, emit) async {
      emit(AuthLoading());
      final result =
          await _loginUsecase(LoginParams(loginEntity: event.loginEntity));
      result.fold((failure) => emit(AuthError(message: failure.message)),
          (userData) => emit(AuthSuccess()));
    });

    on<RegisterEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _registerUsecase(
          RegisterParams(registrationEntity: event.registrationEntity));
      result.fold((failure) => emit(AuthError(message: failure.message)),
          (userData) => emit(AuthRegisterSuccess()));
    });

    on<GetUserEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _getUserUsecase(NoParams());
      result.fold((failure) => emit(AuthError(message: failure.message)),
          (userData) => emit(AuthAuthenticated(userDataEntity: userData)));
    });

    on<LogoutEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _logoutUsecase(NoParams());
      result.fold((failure) => emit(AuthError(message: failure.message)),
          (userData) => emit(AuthLoggedOut()));
    });
  }
}
