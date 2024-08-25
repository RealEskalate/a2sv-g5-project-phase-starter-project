class SellerModel {
  final String id;
  final String name;
  final String email;

  SellerModel({
    required this.id,
    required this.name,
    required this.email,
  });

  factory SellerModel.fromJson(Map<String, dynamic> json) {
    return SellerModel(
      id: json['_id'],
      name: json['name'],
      email: json['email'],
    );
  }
}