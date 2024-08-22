import 'package:equatable/equatable.dart';

class UserEntity extends Equatable{
  final String name;
  final String email;
  final String password;

  const UserEntity({
    required this.name,
    required this.email,
    required this.password,

  });
  
  @override
  
  List<Object?> get props => [name,email,password];

  Map<String, String> toJsonentitiy() {
    return {
      'email': email,
      'password': password,
    };
  }
    
}