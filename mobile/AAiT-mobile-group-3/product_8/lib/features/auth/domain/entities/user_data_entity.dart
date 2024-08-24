import 'package:equatable/equatable.dart';

class UserDataEntity extends Equatable {
  final String id;
  final String email;
  final String name;


  const UserDataEntity({
    required this.id,
    required this.email,
    required this.name,
   
  });

  @override
  List<Object?> get props => [email, name];
}

