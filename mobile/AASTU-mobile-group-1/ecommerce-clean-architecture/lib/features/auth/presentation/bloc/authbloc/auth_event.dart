import 'package:equatable/equatable.dart';

abstract class  UserEvent extends Equatable {
  const UserEvent();
  @override
  List<Object> get props => [];
} 

class LoginEvent extends UserEvent {
  final String email;
  final String password;
  const LoginEvent({required this.email, required this.password});
  @override
  List<Object> get props => [email, password];
}
class RegisterEvent extends UserEvent {
  final String email;
  final String password;
  final String name;
  const RegisterEvent({required this.email, required this.password,required this.name});
  @override
  List<Object> get props => [email, password,name];
}