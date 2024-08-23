import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: ContactsPage(),
    );
  }
}

class ContactsPage extends StatelessWidget {
  final List<Contact> contacts = [
    Contact(name: 'Antonio Banderas', status: 'Online', isOnline: true),
    Contact(name: 'Antonio Banderas', status: 'Online', isOnline: true),
    Contact(name: 'Bessie Cooper', status: 'Last seen today at 8:40'),
    Contact(name: 'Leslie Alexander', status: 'Last seen today at 8:40'),
    Contact(name: 'Jacob Jones', status: 'Last seen today at 8:40'),
    Contact(name: 'Leslie Alexander', status: 'Last seen today at 8:40'),
    Contact(name: 'Floyd Miles', status: 'Last seen long time ago'),
  ];

  ContactsPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: const FloatingActionButton(
        backgroundColor: Colors.black,
        onPressed: null,
        child: Icon(Icons.add_comment),
      ),
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        title: const Text('Contacts', style: TextStyle(color: Colors.black)),
        actions: [
          IconButton(
            icon:const FaIcon(
              FontAwesomeIcons.magnifyingGlass,
              size: 30,
              color: Colors.black
            ),
            onPressed: () {},
          ),
        ],
        // leading: IconButton(
        //   icon: Icon(Icons.arrow_back_ios, color: Colors.black),
        //   onPressed: () {},
        // ),
      ),
      body: Column(
        children: [
          const ListTile(
            leading: Icon(Icons.people, color: Colors.black),
            title: Text('Invite friends'),
            onTap: null,
          ),
          const ListTile(
            leading: Icon(Icons.location_on, color: Colors.black),
            title: Text('Find people nearby'),
            onTap: null,
          ),
          const Divider(),
          Expanded(
            child: ListView.builder(
              itemCount: contacts.length,
              itemBuilder: (context, index) {
                return ContactTile(contact: contacts[index]);
              },
            ),
          ),
        ],
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: 1,
        selectedItemColor: Colors.black,
        unselectedItemColor: Colors.grey,
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.message),
            label: 'Messages',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.contacts),
            label: 'Contacts',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.call),
            label: 'Calls',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
      ),
    );
  }
}

class Contact {
  final String name;
  final String status;
  final bool isOnline;

  Contact({required this.name, required this.status, this.isOnline = false});
}

class ContactTile extends StatelessWidget {
  final Contact contact;

  ContactTile({required this.contact});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: CircleAvatar(
        backgroundColor: Colors.grey,
        child: Text(
          contact.name.substring(0, 2),
          style: const TextStyle(color: Colors.white),
        ),
      ),
      title: Text(contact.name),
      subtitle: Text(
        contact.status,
        style: TextStyle(
          color: contact.isOnline ? Colors.green : Colors.grey,
        ),
      ),
      trailing: contact.isOnline
        ? const Icon(
            Icons.circle,
            color: Colors.green,
            size: 12,
          )
        : null,
    );
  }
}
