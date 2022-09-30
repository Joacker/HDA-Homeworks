
-- CREATE TABLE Users(Id_user SERIAL, Nombre VARCHAR, Apellido VARCHAR, Email VARCHAR, Password VARCHAR);
-- CREATE TABLE Champions(Id_partido INT, Nombre_equipo1 VARCHAR, Nombre_equipo2 VARCHAR, Categoria VARCHAR, Resultado INT);

CREATE TABLE Agents (
  Id SERIAL PRIMARY KEY,
  Name varchar,
  Role varchar,
  Descripcion varchar,
  Skill1 varchar,
  Skill2 varchar,
  Skill3 varchar,
  Ulti varchar,
  Url_image varchar
);

CREATE TABLE Resenias (
  Id int PRIMARY KEY,
  Id_agente int,
  Resenia varchar,
  Id_user int
);

CREATE TABLE Users (
  Id SERIAL PRIMARY KEY,
  Email varchar,
  Password varchar
);

--ALTER TABLE Resenias ADD FOREIGN KEY (Id_agente) REFERENCES Agentes (Id);

--ALTER TABLE Resenias ADD FOREIGN KEY (Id_user) REFERENCES Usuarios (Id);

INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Astra', 'CONTROLLER', 'Ghanaian Agent Astra harnesses the energies of the cosmos to reshape battlefields to her whim. With full command of her astral form and a talent for deep strategic foresight, she is always eons ahead of her enemy is next move.', 'Nova Pulse', 'Nebula / Dissipate', 'Gravity Well', 'Astral Form / Cosmic Divide', 'https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/astra_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Breach', 'INITIATOR', 'The bionic Swede Breach fires powerful, targeted kinetic blasts to aggressively clear a path through enemy ground. The damage and disruption he inflicts ensures no fight is ever fair.', 'Fault Line', 'Flashpoint', 'Rolling Thunder', 'Fault Line / Fault Line','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/breach_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Brimstone', 'CONTROLLER', 'Joining from the USA, Brimstone is orbital arsenal ensures his squad always has the advantage. His ability to deliver utility precisely and safely make him the unmatched boots-on-the-ground commander.', 'Incendiary', 'Sky Smoke', 'Stim Beacon', 'Orbital Strike','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/brimstone_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Cypher', 'SENTINEL', 'The Moroccan information broker, Cypher is a one-man surveillance network who keeps tabs on the enemy is every move. No secret is safe. No maneuver goes unseen. Cypher is always watching.', 'Cyber Cage', 'Spycam', 'Trapwire', 'Neural Theft','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/cypher_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Jett', 'DUELIST', 'Representing her home country of South Korea, Jett is agile and evasive fighting style lets her take risks no one else can. She runs circles around every skirmish, cutting enemies up before they even know what hit them.', 'Updraft', 'Tailwind', 'Cloudburst', 'Blade Storm','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/jett_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Killjoy', 'SENTINEL', 'The genius of Germany, Killjoy secures and defends key battlefield positions with a collection of traps, turrets, and mines. Each invention is primed to punish any assailant too dumb to back down.','Nanoswarm', 'ALARMBOT', 'TURRET', 'Lockdown','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/killjoy_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Omen', 'CONTROLLER', 'A phantom of a memory, Omen hunts in the shadows. He renders enemies blind, teleports across the field, then lets paranoia take hold as his foe scrambles to uncover where he might strike next.','Paranoia', 'Dark Cover', 'Shrouded Step', 'From the Shadows','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/omen_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Phoenix', 'DUELIST', 'Hailing from the U.K., Phoenix is star power shines through in his fighting style, igniting the battlefield with flash and flare. Whether he is got backup or not, he is rushing in to fight on his own terms.', 'Curveball', 'Hot Hands', 'Blaze', 'Run it Back','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/phoenix_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Raze', 'DUELIST', 'Raze explodes out of Brazil with her big personality and big guns. With her blunt-force-trauma playstyle, she excels at flushing entrenched enemies and clearing tight spaces with a generous dose of "boom".', 'Boom Bot', 'Blast Pack', 'Paint Shells', 'Showstopper','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/raze_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Reyna', 'DUELIST', 'Forged in the heart of Mexico, Reyna dominates single combat, popping off with each kill she scores. Her capability is only limited by her raw skill, making her sharply dependant on performance.', 'Leer', 'Dispersion', 'Empress', 'Death Blossom','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/reyna_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Sage', 'SENTINEL', 'The bastion of China, Sage creates safety for herself and her team wherever she goes. Able to revive fallen friends and stave off forceful assaults, she provides a calm center to a hellish battlefield.' , 'Slow Orb', 'Healing Orb', 'Barrier Orb', 'Resurrection','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/sage_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Skye', 'INITIATOR', 'Hailing from Australia, Skye and her band of beasts trailblaze the way through hostile territory. With her creations hampering the enemy, and her power to heal others, the team is strongest and safest by Skye is side.','Trailblazer', 'Guiding Light', 'Blinding Orb', 'Star Sheriff','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/skye_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Sova', 'INITIATOR', 'Born from the eternal winter of Russia is tundra, Sova tracks, finds, and eliminates enemies with ruthless efficiency and precision. His custom bow and incredible scouting abilities ensure that even if you run, you cannot hide.', 'Shock Bolt', 'Recon Bolt', 'Owl Drone', 'Hunter’s Fury','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/sova_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Viper', 'CONTROLLER', 'The American Chemist, Viper deploys an array of poisonous chemical devices to control the battlefield and cripple the enemy is vision. If the toxins do not kill her prey, her mindgames surely will.','Toxic Screen', 'Poison Cloud', 'Snake Bite', 'Viper’s Pit','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/viper_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Yoru', 'DUELIST', 'Japanese native Yoru rips holes straight through reality to infiltrate enemy lines unseen. Using deception and aggression in equal measure, he gets the drop on each target before they know where to look.', 'Dimensional Drift', 'Unstable Orb', 'Gatecrash', 'Dark Cover','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/yoru_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('KAY/O', 'INITIATOR', 'KAY/O is a machine of war built for a single purpose: neutralizing radiants. His power to suppress enemy abilities cripples his opponents capacity to fight back, securing him and his allies the ultimate edge.', 'Overdrive', 'Disassemble', 'Reassemble', 'Reassemble / Reassemble','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/kayo_portrait.png');
INSERT INTO Agents(Id, Name, Role, Skill_1, Skill_2, Skill_3, Ulti, Url_image) VALUES('Chamber', 'SENTINEL', 'Well dressed and well armed, French weapons designer Chamber expels aggressors with deadly precision. He leverages his custom arsenal to hold the line and pick off enemies from afar, with a contingency built for every plan.','Rendezvous', 'Trademark', 'Headhunter', 'Tour De Force','https://trackercdn.com/cdn/tracker.gg/valorant/db/agents/chamber_portrait.png');

INSERT INTO Users(Id, Email, Password) VALUES ('Juan', 'Perez', 'juan.perez@gmail.com', '1234');
INSERT INTO Users(Id, Email, Password) VALUES ('Maria', 'Gomez', 'maria.gomez@gmail.com', '1234');
INSERT INTO Users(Id, Email, Password) VALUES ('Pedro', 'Rodriguez', 'pedro.rodriguez@hotmail.com', '1234');
INSERT INTO Users(Id, Email, Password) VALUES ('Jose', 'Lopez', 'jose.lopez@yahoo.com', '1234');
INSERT INTO Users(Id, Email, Password) VALUES ('Ana', 'Martinez', 'ana.martinez@hotmail.com', '1234');