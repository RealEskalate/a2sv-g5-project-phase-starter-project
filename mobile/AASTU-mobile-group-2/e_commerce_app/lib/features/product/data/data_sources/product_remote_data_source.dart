import 'dart:convert';
import 'dart:io';

import 'package:e_commerce_app/core/constants/constants.dart';
import 'package:e_commerce_app/features/product/data/models/product_model.dart';
import 'package:e_commerce_app/features/product/domain/usecase/insert_product_usecase.dart';
import 'package:equatable/equatable.dart';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/failure/exception.dart';
import 'dart:io';

// void main() async {
//   ProductModel testProduct = ProductModel(
//       description: "not 2.0",
//       id: "6672752cbd218790438efdb0",
//       imageUrl:
//           "D:/abd/A2SV/2024-internship-mobile-tasks/e_commerce_app/assets/download.jpeg",
//       name: "Anime website",
//       price: 123);
//   http.Client client = http.Client();
//   ProductRemoteDataSource a = ProductRemoteDataSource(client: client);
//   // a.insertProduct(testProduct).then((product) {
//   //   print("Product inserted: ${product.name}");
//   // }).catchError((error) {
//   //   print("Failed to insert product: $error");
//   // });

//   // List<ProductModel> products = await a.getAllProduct();
//   // print(products);
// }

abstract class ProductDataSource extends Equatable {
  Future<List<ProductModel>> getAllProduct();
  Future<ProductModel> getOneProduct(String id);
  Future<ProductModel> insertProduct(ProductModel newProduct);
  Future<ProductModel> updateProduct(ProductModel updatedProduct);
  Future<String> deleteProduct(String id);

  @override
  // TODO: implement props
  List<Object?> get props => [];
}

class ProductRemoteDataSource extends ProductDataSource {
  final http.Client client;
  ProductRemoteDataSource({required this.client});

  @override
  Future<List<ProductModel>> getAllProduct() async {
    var accessToken = await getToken();
    final response = await client.get(Uri.parse(Urls.baseUrl), headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $accessToken',
    });
    print(response.statusCode);
    print(response.body);
    if (response.statusCode == 200) {
      return ProductModel.fromJsonList(json.decode(response.body)["data"]);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<ProductModel> getOneProduct(String id) async {
    var accessToken = await getToken();
    final response =
        await client.get(Uri.parse(Urls.getProductById(id)), headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $accessToken',
    });

    if (response.statusCode == 200) {
      return ProductModel.fromJson(json.decode(response.body)["data"]);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<ProductModel> insertProduct(ProductModel newProduct) async {
    var accessToken = await getToken();
    File image = File(newProduct.imageUrl);

    if (!image.existsSync()) {
      // print("image doesnt exist");
      throw ServerException();
    } else {
      print("image found");
    }
    print("image found");

    var request = http.MultipartRequest(
        'POST',
        Uri.parse(
          Urls.baseUrl,
        ));
    request.files.add(
      await http.MultipartFile.fromPath('image', image.path,
          contentType: MediaType('image', "jpeg")),
    );
    request.fields["name"] = newProduct.name;
    request.fields["description"] = newProduct.description;
    request.fields["price"] = newProduct.price.toString();
    request.headers.addAll({
      'Content-Type': 'multipart/form-data',
      'Authorization': 'Bearer $accessToken',
    });
    final response = await request.send();

    final responseBody = await response.stream.bytesToString();
print("Response Body: $responseBody");

    if (response.statusCode == 201 || response.statusCode == 200) {
      // print("Uploaded successfully");
      return ProductModel.fromJson(json.decode(responseBody));
    } else {
      throw ServerException();
    }
  }

  @override
  Future<ProductModel> updateProduct(ProductModel updatedProduct) async {
    var accessToken = await getToken();
    try {
      var cd = Urls.getProductById(updatedProduct.id);

      var ab = jsonEncode(updatedProduct.toJson());

      final response = await client.put(Uri.parse(cd), body: ab, headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      });

      if (response.statusCode == 200) {
        return ProductModel.fromJson(json.decode(response.body));
      } else {
        throw ServerException();
      }
    } catch (e) {
      throw ServerException();
    }
  }

  @override
  Future<String> deleteProduct(String id) async {
    var accessToken = await getToken();
    final response =
        await client.delete(Uri.parse(Urls.getProductById(id)), headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $accessToken',
    });

    if (response.statusCode == 200) {
      return "deleted";
    } else {
      throw ServerException();
    }
  }

  static Future<String> getToken() async {
    var sharedPreferences = await SharedPreferences.getInstance();

    final userToken = sharedPreferences.getString("user");
    print(userToken);
    if (userToken != null) {
      return userToken;
    } else {
      throw Exception("no user have logged in");
    }
  }
}
