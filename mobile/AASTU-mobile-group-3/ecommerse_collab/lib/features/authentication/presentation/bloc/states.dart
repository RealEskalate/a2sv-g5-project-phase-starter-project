import '../../domain/entity/user.dart';

class UserState {}

class UserInitialState extends UserState {}

class UserAuthenticatedState extends UserState{
  final User user;
  UserAuthenticatedState(this.user);

}

class LoginLoadingState extends UserState {}

class RegisterLoadingState extends UserState {}

class LoginErrorState extends UserState {
  final String message;
  LoginErrorState(this.message);
} 
class RegisterErrorState extends UserState {
  final String message;
  RegisterErrorState(this.message);

}

class UserRegisteredState extends UserState{
  final User user;
  UserRegisteredState(this.user);
}

class UserLoggedState extends UserState{
  final User user;
  
  UserLoggedState(this.user);
}


class LoggedOutState extends UserState{
  
}
class LogOutLoadingState extends UserState{
  
}