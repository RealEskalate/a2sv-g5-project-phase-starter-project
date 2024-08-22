import 'package:equatable/equatable.dart';

class UserDataEntity extends Equatable {
  final String email;
  final String name;


  const UserDataEntity({
    required this.email,
    required this.name,
   
  });

  @override
  List<Object?> get props => [email, name];
}

