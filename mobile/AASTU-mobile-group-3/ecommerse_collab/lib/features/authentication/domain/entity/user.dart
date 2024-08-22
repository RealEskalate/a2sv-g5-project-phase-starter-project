import 'package:equatable/equatable.dart';

class User extends Equatable{
  final String id;
  final String username;
  final String password;
  final String email;

  User({ required this.id ,required this.username,required this.password, required this.email });
  
  @override
  
  List<Object?> get props => [id, username, password, email];


}