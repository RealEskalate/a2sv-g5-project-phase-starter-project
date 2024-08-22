import 'package:equatable/equatable.dart';

class UserDataEntity extends Equatable {
  final String name;
  final String email;

  UserDataEntity({required this.email, required this.name});

  @override
  List<Object?> get props => [email, name];
}
