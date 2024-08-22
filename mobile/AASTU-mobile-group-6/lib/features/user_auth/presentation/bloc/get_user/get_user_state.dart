import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:equatable/equatable.dart';


abstract class GetUserState extends Equatable {
  const GetUserState();
}



class GetUserLoading extends GetUserState {
  @override
  List<Object> get props => [];
}

class GetUserLoaded extends GetUserState {
    final UserModel user;
    

  const GetUserLoaded(this.user);

  @override
  List<Object> get props => [user];
}

class GetUserFailure extends GetUserState {
  final String error;

  const GetUserFailure(this.error);

  @override
  List<Object> get props => [error];
}