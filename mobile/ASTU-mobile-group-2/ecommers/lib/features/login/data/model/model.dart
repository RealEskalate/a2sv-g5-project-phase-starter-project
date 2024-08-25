



import '../../domain/entity/login_entity.dart';

class LoginModel extends LoginEntity {

  LoginModel({
    required super.accessToken, 
    required super.id, 
    required super.name, 
    required super.email, 
     required super.message
    
  });


  factory LoginModel.fromJson(Map<String, dynamic> json) {
    return LoginModel(
      accessToken: json['data']['accessToken'],
      id: json['data']['id'],
      name: json['data']['name'],
      email: json['data']['email'],
      message: json['message']


    );
  }

  LoginEntity toEntity() => LoginEntity(
      accessToken: accessToken,
      id: id,
      name: name,
      email: email,
      message: message

  );
    
}
