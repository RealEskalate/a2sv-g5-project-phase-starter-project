import 'package:bloc/bloc.dart';

import '../../domain/usecases/log_out.dart';
import '../../domain/usecases/login.dart';
import '../../domain/usecases/sign_up.dart';
import 'events.dart';
import 'states.dart';

class UserBloc extends Bloc<UserEvent, UserState>{
  final SignUpUseCase signUpUseCase;
  final LoginUseCase loginUseCase;
  final LogOutUseCase logOutUseCase;
  UserBloc({
    required this.logOutUseCase,
    required this.signUpUseCase,
    required this.loginUseCase
  }):super(UserInitialState()){
    on<RegisterUserEvent>((event, emit) async {
      emit (RegisterLoadingState());
      try{
        final user = await signUpUseCase.call(email: event.email, password: event.password, username: event.username);
        // final user =  User(id: '1', email: event.email, password: event.password, username: event.username);
        print("from user bloc $user");
        emit(UserRegisteredState(user));
      } catch (e) {
          print('Error during registration: $e');
          emit(RegisterErrorState("Fail to Register"));
      }

    },);
    on<LogInEvent> ((event, emit) async {
      emit(LoginLoadingState());
      try{
        final user = await loginUseCase.call( event.email,event.password);
        user.fold((l) => emit(LoginErrorState("Fail to Login")), (r) => emit(UserLoggedState(r)));  
      } catch (e) {
        emit(LoginErrorState("Fail to Login")); 
     }
      
    });

    on<LogOutEvent> ((event, emit)async {
      emit(LogOutLoadingState());
      try{
        await logOutUseCase.call();
        emit(LoggedOutState());
      }catch (e){
        throw e;
      }

    });

  }



  
  
}