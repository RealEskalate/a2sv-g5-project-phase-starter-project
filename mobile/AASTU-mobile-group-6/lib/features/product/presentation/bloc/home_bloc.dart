import 'package:bloc/bloc.dart';
import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/local_data_source/local_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/repositories/product_repository_impl.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_state.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;

import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:meta/meta.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_detail_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/delete_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/update_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/add_usecase.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/usecases/usecases.dart';
import '../../domain/usecases/get_all_usecase.dart';





class HomeBloc extends Bloc<HomeEvent, HomeState>  {
  final client = http.Client;
  final localData = SharedPreferences.getInstance();
  GetAllUsecase getAllProductsUseCase;
  GetDetailUseCase getDetailUsecase;

  HomeBloc(this.getAllProductsUseCase,this.getDetailUsecase) : super(HomeLoading()) {
    on<GetProductsEvent>((event, emit) async{
      emit(HomeLoading());  
      var products = await getAllProductsUseCase(NoParams());
      products.fold((l) => emit(HomeFailure(l.message)), (r) => emit(HomeLoaded(r)));

    });
    
  }
}
