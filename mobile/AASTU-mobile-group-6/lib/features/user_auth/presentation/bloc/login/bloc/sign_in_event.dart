import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:equatable/equatable.dart';


abstract class LoginEvent extends Equatable {
  const LoginEvent();
}

class LogUserIn extends LoginEvent {
  final UserModel user;
  LogUserIn(this.user);
  @override
  List<Object> get props => [user];
}