import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:equatable/equatable.dart';


abstract class GetUserEvent extends Equatable {
  const GetUserEvent();
}

class GetUserInfoEvent extends GetUserEvent {
  
  @override
  List<Object> get props => [];
}