import 'package:ecommerce/features/auth/domain/repository/user_repository.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_event.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_state.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../domain/usecases/loginUser.dart';
import '../../../domain/usecases/registerUser.dart';

class UserBloc extends Bloc<UserEvent, UserState> {
  final RegisterUserUseCase registerUserUsecase;
  final LoginUserUsecase loginUserUsecase;
  UserBloc({required this.loginUserUsecase,required this.registerUserUsecase}) : super(registering())
  {
    on<RegisterEvent>(_onRegister);
    on<LoginEvent>(_onLogin);
  }


Future<void>_onRegister(RegisterEvent event, Emitter<UserState> emit)async{
  emit(registering());

   var user = await registerUserUsecase.register(event.name, event.email, event.password);
    print('from bloc');
  print(user);
    user.fold(
      (failure) => emit(registerfailure(failure.message)),
      (_) => emit(registered()),
    );

}
Future<void>_onLogin(LoginEvent event, Emitter<UserState> emit)async{
  emit(logging());

   var user = await loginUserUsecase.login(event.email, event.password);
    print('user ${user}');
    user.fold(
      (failure) => emit(logginfailure(failure.message)),
      (userModel) => emit(logged(userModel)),
    );

}
}
