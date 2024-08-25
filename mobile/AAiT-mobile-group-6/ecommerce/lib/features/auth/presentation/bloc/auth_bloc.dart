import 'package:bloc/bloc.dart';

import 'package:meta/meta.dart';

import '../../../product/domain/usecase/get_all_product.dart';
import '../../domain/entities/log_in_entity.dart';
import '../../domain/entities/sign_up_entity.dart';
import '../../domain/usecase/log_out_usecase.dart';
import '../../domain/usecase/login_usecase.dart';
import '../../domain/usecase/signUp_usecase.dart';

part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final SignUpUseCase signUpUseCase;
  final LogInUseCase logInUseCase;
  final LogoutUsecase loguoutUsecase;

  AuthBloc(this.signUpUseCase, this.logInUseCase,this.loguoutUsecase) : super(AuthInitial()) {
    on<SingUpEvent>((event, emit) async {
      emit(AuthLoadingState());

      final res = await signUpUseCase(UseCaseParams(event.signUpEntity));

      res.fold(
        (l) {
          emit(AuthErrorState(message: l.message));
        },
        (r) {
          emit(AuthSuccessState(message: r));
        },
      );
    });

    on<LogInEvent>((event, emit) async {
      emit(AuthLoadingState());
      final res = await logInUseCase(LogInParams(event.logInEntity));

      res.fold(
        (l) {
          emit(AuthErrorState(message: l.message));
        },
        (r) {
          emit(AuthSuccessState(message: r));
        },
      );
    });

    on<LogOutEvent>(
      (event, emit) async {
        emit(AuthLoadingState());
        final result = await loguoutUsecase(NoParams());
        result.fold(
          (failure) => emit(UserLogoutState(
              message: "Failed to logout, please try again",
              status: AuthStatus.error)),
          (response) => emit(UserLogoutState(
              message: "Logged out successfully", status: AuthStatus.loaded)),
        );
      },
    );
  }
}
