import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../core/Error/failure.dart';
import '../../../../core/const/const.dart';
import '../model/ecommerce_model.dart';

/// Abstract class representing the remote data source for the Ecommerce feature.
abstract class EcommerceRemoteDataSource {
  
  /// Retrieves a single product by its [id].
  Future<EcommerceModel> getProduct(String id);

  /// Retrieves all products.
  Future<List<EcommerceModel>> getAllProducts();

  /// Edits a product identified by its [id] with the provided [model].
  Future<bool> editProduct(String id, Map<String,dynamic> data);

  /// Deletes a product identified by its [id].
  Future<bool> deleteProduct(String id);

  /// Adds a new product with the provided [data].
  Future<bool> addProduct(Map<String,dynamic> data);
}

/// Implementation of the [EcommerceRemoteDataSource] interface.
class EcommerceRemoteDataSourceImpl implements EcommerceRemoteDataSource {
  final http.Client client;
  final SharedPreferences sharedPreferences;

  EcommerceRemoteDataSourceImpl({
    required this.client,
    required this.sharedPreferences
  });
  
  @override
  Future<EcommerceModel> getProduct(String id) async {
    final token = sharedPreferences.getString('key');

      final response = await client.get(
        Uri.parse(Urls.getByUrl(id)),
        headers: {
          'Authorization': 'Bearer $token',
        },
      );

    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      if (data != null) {
   
        return EcommerceModel.fromJson(data['data']);
      } else {
        throw Exception('Error no data source');
      }
    } else {
      throw const ConnectionFailur(message: 'server error');
    }
  }

  @override
  Future<List<EcommerceModel>> getAllProducts() async {
    try {
      
      final token = sharedPreferences.getString('key');
    
      final response = await client.get(
        Uri.parse(Urls.getAll()),
        headers: {
          'Authorization': 'Bearer $token',
        },
      );
      if (response.statusCode == 200) {
        final data = json.decode(response.body);
     
        if (data != null) {

          return EcommerceModel.getAllProduct(data);
        } else {
          throw const ServerFailure(message: 'server Error');
        }
      } else {
        throw const ConnectionFailur(message: 'server Error');
      } } catch (e) {
       
        throw ConnectionFailur(message: e.toString());
      }
  }
  
  @override
  Future<bool> deleteProduct(String id) async {
    final token = sharedPreferences.getString('key');

      final response = await client.delete(
        Uri.parse(Urls.deleteProduct(id)),
        headers: {
          'Authorization': 'Bearer $token',
        },
      );
    if (response.statusCode == 200) {
      return Future.value(true);
    } else {
      return Future.value(false);
    }
  }
  
  @override
  Future<bool> editProduct(String id, Map<String, dynamic> data) async {
   final token = sharedPreferences.getString('key');
  final response = await client.put(
    Uri.parse(Urls.updateProduct(id)),
    headers: {'Content-Type': 'application/json','Authorization': 'Bearer $token',},
    body: jsonEncode({
      'name': data['name'],
      'description': data['description'],
      'price': data['price'],
    }),
    
  );  // Log the response body for detailed error messages

  if (response.statusCode == 200) {
    return Future.value(true);
  } else {
    return Future.value(false);
  }
}

@override  
 Future<bool> addProduct(Map<String, dynamic> data) async {
  try {
    // Debugging output
   
    
   final token = sharedPreferences.getString('key');
    
    var request = http.MultipartRequest('POST', Uri.parse(Urls.addNewProduct()))
      ..fields['name'] = data['name']
      ..fields['price'] = data['price']
      ..fields['description'] = data['description']
      ..headers['Authorization'] = 'Bearer $token'
      ..files.add(await http.MultipartFile.fromPath(
        'image', 
        data['file'].path,
        contentType: MediaType('image', 'png')),
        
        );
    
  
    
    var response = await request.send();
    
    // Read the response
  

    
    if (response.statusCode == 201) {
      return Future.value(true);
    } else {
      return Future.value(false);
    }
  } catch (e) {
    return Future.value(false);
  }
}
}