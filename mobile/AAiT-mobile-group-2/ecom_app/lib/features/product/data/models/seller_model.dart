import '../../../../core/error/exception.dart';
import '../../domain/entities/seller.dart';

class SellerModel extends Seller {
  const SellerModel({
    required super.id,
    required super.name,
    required super.email,
  });

  factory SellerModel.fromJson(Map<String, dynamic> json) {
    try {
      return SellerModel(
        id: json['_id'],
        name: json['name'],
        email: json['email'],
      );
    } catch (e) {
      throw JsonParsingException();
    }
  }
  static List<SellerModel> fromJsonList(List<dynamic> jsonList) {
    try {
      return jsonList
          .map((json) => SellerModel.fromJson(json as Map<String, dynamic>))
          .toList();
    } catch (e) {
      throw JsonParsingException();
    }
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'name': name,
      'email': email,
    };
  }

  static List<Map<String, dynamic>> toJsonList(List<SellerModel> products) {
    return products.map((product) => product.toJson()).toList();
  }

  Seller toEntity() => Seller(
        id: id,
        name: name,
        email: email,
      );

  static List<Seller> toEntityList(List<SellerModel> models) {
    return models.map((model) => model.toEntity()).toList();
  }

  static Seller toModel(Seller seller) {
    return SellerModel(
      id: seller.id,
      name: seller.name,
      email: seller.email,
    );
  }
}
